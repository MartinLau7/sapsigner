package emulator

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/log"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/interposer"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/trace"
)

const (
	pageSize        = uint64(0x1000)
	startAddr       = uint64(0x0000000100000000)
	finishAddr      = startAddr - 1
	interposerAddr  = uint64(0x0000000100004000)
	bridgeAddr      = uint64(0x0000600000000000)
	heapHeadAddr    = uint64(0x00007FF7B0000000)
	heapTailAddr    = uint64(0x00007FF7B4000000)
	heapSize        = heapTailAddr - heapHeadAddr
	stackTopAddr    = uint64(0x00007FF7BF800000)
	stackBottomAddr = uint64(0x00007FF7C0000000)
	stackSize       = stackBottomAddr - stackTopAddr
	mescalAddr      = uint64(0x00007FF800000000)

	// startAddr      = uint64(0x00007FF800000000)
	// interposerAddr = uint64(0x00007FF800004000)
	// mescalAddr     = uint64(0x0000000100000000)
)

var (
	ip = interposer.NewInterposer(interposerAddr)
)

type Emulator struct {
	mescalO *library.Object
	unicorn unicorn.Unicorn
	pageSet map[uint64]struct{}
}

func NewEmulator(mo *library.Object) (*Emulator, error) {
	if err := mo.Fixup(mescalAddr, ip.SymbolTable()); err != nil {
		return nil, err
	}

	uc, err := unicorn.NewUnicorn(unicorn.ARCH_X86, unicorn.MODE_64)
	if err != nil {
		return nil, err
	}

	e := Emulator{
		mescalO: mo,
		unicorn: uc,
		pageSet: make(map[uint64]struct{}),
	}

	if err := e.hookRDTSC(); err != nil {
		return nil, err
	}

	if err := e.loadInterposer(); err != nil {
		return nil, err
	}

	if err := e.loadMescal(); err != nil {
		return nil, err
	}

	if err := e.setupHeap(); err != nil {
		return nil, err
	}

	if err := e.setupStack(); err != nil {
		return nil, err
	}

	if err := e.setupStart(); err != nil {
		return nil, err
	}

	if err := e.setupTrace(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Emulator) DumpHeap() ([]byte, error) {
	symbolTable := ip.SymbolTable()

	heapHead, err := e.unicorn.MemRead(symbolTable["_malloc.heap_head"], 8)
	if err != nil {
		return nil, err
	}
	heapHeadAddr := binary.LittleEndian.Uint64(heapHead)

	heapTail, err := e.unicorn.MemRead(symbolTable["_malloc.heap_tail"], 8)
	if err != nil {
		return nil, err
	}
	heapTailAddr := binary.LittleEndian.Uint64(heapTail)

	heapSize := heapTailAddr - heapHeadAddr

	return e.unicorn.MemRead(heapHeadAddr, heapSize)
}

func (e *Emulator) hookRDTSC() error {
	callback := func(uc unicorn.Unicorn) bool {
		if err := uc.RegWrite(unicorn.X86_REG_RDX, 0); err != nil {
			panic(err)
		}

		if err := uc.RegWrite(unicorn.X86_REG_RAX, 0); err != nil {
			panic(err)
		}

		return true
	}

	if _, err := e.unicorn.HookAdd(unicorn.HOOK_INSN, callback, 1, 0, unicorn.X86_INS_RDTSC); err != nil {
		logger := log.Logger()
		logger.Printf("libunicorn version does not support hooking RDTSC instructions: %s", err.Error())
	}

	return nil
}

func (e *Emulator) loadInterposer() error {
	for _, section := range ip.Sections() {
		sectionAddr := section.Addr()
		mappingAddr := sectionAddr
		mappingAddr -= (pageSize + mappingAddr) % pageSize

		sectionSize := section.Size()
		mappingSize := sectionSize
		mappingSize += sectionAddr - mappingAddr
		mappingSize += (pageSize - mappingSize) % pageSize

		for i := uint64(0); i < mappingSize; i += pageSize {
			addr := mappingAddr + i

			if _, ok := e.pageSet[addr]; ok {
				continue
			}
			e.pageSet[addr] = struct{}{}

			if err := e.unicorn.MemMap(mappingAddr, pageSize); err != nil {
				return err
			}
		}

		sectionData := section.Data()
		if err := e.unicorn.MemWrite(sectionAddr, sectionData); err != nil {
			return err
		}
	}

	return nil
}

func (e *Emulator) loadMescal() error {
	size := e.mescalO.Size()
	size += (pageSize - size) % pageSize

	if err := e.unicorn.MemMap(mescalAddr, size); err != nil {
		return err
	}

	data := e.mescalO.Data()

	return e.unicorn.MemWrite(mescalAddr, data)
}

func (e *Emulator) setupHeap() error {
	if err := e.unicorn.MemMap(heapHeadAddr, heapSize); err != nil {
		return err
	}

	var addrData [8]byte
	binary.LittleEndian.PutUint64(addrData[:], heapHeadAddr)

	var sizeData [8]byte
	binary.LittleEndian.PutUint64(sizeData[:], heapSize)

	symbolTable := ip.SymbolTable()

	if err := e.unicorn.MemWrite(symbolTable["_malloc.heap_head"], addrData[:]); err != nil {
		return err
	}

	if err := e.unicorn.MemWrite(symbolTable["_malloc.heap_tail"], addrData[:]); err != nil {
		return err
	}

	if err := e.unicorn.MemWrite(symbolTable["_malloc.heap_size"], sizeData[:]); err != nil {
		return err
	}

	return nil
}

func (e *Emulator) setupStack() error {
	if err := e.unicorn.MemMap(stackTopAddr, stackSize); err != nil {
		return err
	}

	return e.unicorn.RegWrite(unicorn.X86_REG_RSP, stackBottomAddr)
}

func (e *Emulator) setupStart() error {
	if err := e.unicorn.MemMap(startAddr, pageSize); err != nil {
		return err
	}

	data := []byte{
		0xFF, 0xD0, // callq *%rax
		0xF4, // hlt
	}
	return e.unicorn.MemWrite(startAddr, data)
}

func (e *Emulator) setupTrace() error {
	logger := log.Logger()

	//goland:noinspection GoBoolExpressions
	if trace.Flags&trace.FlagHookCode != 0 {
		hookCode := func(uc unicorn.Unicorn, addr uint64, size uint32) {
			var instruction []byte
			if trace.Flags&trace.FlagOnCodeHookDisassemble != 0 {
				bytes, err := uc.MemRead(addr, uint64(size))
				if err != nil {
					logger.Print(err)
				}

				instruction = bytes
			}

			_, _ = fmt.Fprintf(trace.Output, "->  0x%x <+?>: ", addr)
			for _, b := range instruction {
				_, _ = fmt.Fprintf(trace.Output, "%02x ", b)
			}
			_, _ = fmt.Fprintf(trace.Output, "\n")
		}
		if _, err := e.unicorn.HookAdd(unicorn.HOOK_CODE, hookCode, 1, 0); err != nil {
			return err
		}
	}

	//goland:noinspection GoBoolExpressions
	if trace.Flags&trace.FlagHookMemValid != 0 {
		hookMemValid := func(uc unicorn.Unicorn, access int, addr uint64, size int, value int64) {
			var data []byte
			if trace.Flags&trace.FlagOnValidMemHookRead != 0 {
				bytes, err := uc.MemRead(addr, uint64(size))
				if err != nil {
					logger.Print(err)
				}

				data = bytes
			}

			switch access {
			case unicorn.MEM_FETCH:
				_, _ = fmt.Fprintf(trace.Output, "F.")
			case unicorn.MEM_READ:
				_, _ = fmt.Fprintf(trace.Output, "R.")
			case unicorn.MEM_WRITE:
				_, _ = fmt.Fprintf(trace.Output, "W.")
			default:
				_, _ = fmt.Fprintf(trace.Output, "% 2d", access)
			}
			_, _ = fmt.Fprintf(trace.Output, "  0x%x [%0 2d]: ", addr, size)
			for _, b := range data {
				_, _ = fmt.Fprintf(trace.Output, "%02x ", b)
			}
			_, _ = fmt.Fprintf(trace.Output, "\n")
		}
		if _, err := e.unicorn.HookAdd(unicorn.HOOK_MEM_VALID&^unicorn.HOOK_MEM_FETCH, hookMemValid, 1, 0); err != nil {
			return err
		}
	}

	//goland:noinspection GoBoolExpressions
	if trace.Flags&trace.FlagHookMemInvalid != 0 {
		hookMemInvalid := func(uc unicorn.Unicorn, access int, addr uint64, size int, value int64) bool {
			var data []byte
			if trace.Flags&trace.FlagOnInvalidMemHookRead != 0 {
				bytes, err := uc.MemRead(addr, uint64(size))
				if err != nil {
					logger.Print(err)
				}

				data = bytes
			}

			switch access {
			case unicorn.MEM_FETCH_UNMAPPED:
				_, _ = fmt.Fprintf(trace.Output, "F?")
			case unicorn.MEM_READ_UNMAPPED:
				_, _ = fmt.Fprintf(trace.Output, "R?")
			case unicorn.MEM_WRITE_UNMAPPED:
				_, _ = fmt.Fprintf(trace.Output, "W?")
			case unicorn.MEM_FETCH_PROT:
				_, _ = fmt.Fprintf(trace.Output, "F!")
			case unicorn.MEM_READ_PROT:
				_, _ = fmt.Fprintf(trace.Output, "R!")
			case unicorn.MEM_WRITE_PROT:
				_, _ = fmt.Fprintf(trace.Output, "W!")
			default:
				_, _ = fmt.Fprintf(trace.Output, "% 2d", access)
			}
			_, _ = fmt.Fprintf(trace.Output, "  0x%x [%0 2d]: ", addr, size)
			for _, b := range data {
				_, _ = fmt.Fprintf(trace.Output, "%02x ", b)
			}
			_, _ = fmt.Fprintf(trace.Output, "\n")

			return false
		}
		if _, err := e.unicorn.HookAdd(unicorn.HOOK_MEM_INVALID, hookMemInvalid, 1, 0); err != nil {
			return err
		}
	}

	return nil
}

func (e *Emulator) pushArgs(args ...uint64) error {
	rsp, err := e.unicorn.RegRead(unicorn.X86_REG_RSP)
	if err != nil {
		return err
	}

	for i := len(args) - 1; i >= 0; i-- {
		arg := args[i]

		var err error
		switch i {
		case 0:
			err = e.unicorn.RegWrite(unicorn.X86_REG_RDI, arg)
		case 1:
			err = e.unicorn.RegWrite(unicorn.X86_REG_RSI, arg)
		case 2:
			err = e.unicorn.RegWrite(unicorn.X86_REG_RDX, arg)
		case 3:
			err = e.unicorn.RegWrite(unicorn.X86_REG_RCX, arg)
		case 4:
			err = e.unicorn.RegWrite(unicorn.X86_REG_R8, arg)
		case 5:
			err = e.unicorn.RegWrite(unicorn.X86_REG_R9, arg)
		default:
			err = e.pushArgToStack(&rsp, arg)
		}
		if err != nil {
			return err
		}
	}

	return e.unicorn.RegWrite(unicorn.X86_REG_RSP, rsp)
}

func (e *Emulator) pushArgToStack(rsp *uint64, arg uint64) error {
	var data [8]byte
	binary.LittleEndian.PutUint64(data[:], arg)

	*rsp -= 8

	if err := e.unicorn.MemWrite(*rsp, data[:]); err != nil {
		return err
	}

	return nil
}

var _ io.Closer = (*Emulator)(nil)

func (e *Emulator) Close() error {
	return e.unicorn.Close()
}
