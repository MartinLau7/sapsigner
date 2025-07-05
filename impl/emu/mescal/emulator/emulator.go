package emulator

import (
	"encoding/binary"
	"fmt"
	"io"
	"maps"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/log"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/interposer"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/trace"
)

const (
	pageSize         = uint64(0x1000)
	startAddr        = uint64(0x0000000100000000)
	finishAddr       = startAddr - 1
	interposerAddr   = uint64(0x0000000100004000)
	bridgeAddr       = uint64(0x0000600000000000)
	heapHeadAddr     = uint64(0x00007FF7B0000000)
	heapTailAddr     = uint64(0x00007FF7B4000000)
	heapSize         = heapTailAddr - heapHeadAddr
	stackTopAddr     = uint64(0x00007FF7BF800000)
	stackBottomAddr  = uint64(0x00007FF7C0000000)
	stackSize        = stackBottomAddr - stackTopAddr
	corefpicxsAddr   = uint64(0x00007FF800000000)
	corefpAddr       = uint64(0x00007FF804000000)
	commercecoreAddr = uint64(0x00007FF808000000)
	commercekitAddr  = uint64(0x00007FF80C000000)
	storeagentAddr   = uint64(0x00007FF810000000)
)

type Emulator struct {
	corefpicxsO   *library.Object
	corefpO       *library.Object
	commercecoreO *library.Object
	commercekitO  *library.Object
	storeagentO   *library.Object
	symbolTable   map[string]uint64
	unicorn       unicorn.Unicorn
}

func NewEmulator(corefpicxsO *library.Object, corefpO *library.Object, commercecoreO *library.Object, commercekitO *library.Object, storeagentO *library.Object) (*Emulator, error) {
	ip := interposer.NewInterposer(interposerAddr)
	symbolTable := maps.Clone(ip.SymbolTable())

	if err := corefpicxsO.Fixup(corefpicxsAddr, symbolTable); err != nil {
		return nil, fmt.Errorf("failed to fixup CoreFP.icxs: %w", err)
	}
	maps.Insert(symbolTable, maps.All(corefpicxsO.SymbolTable()))

	if err := corefpO.Fixup(corefpAddr, symbolTable); err != nil {
		return nil, fmt.Errorf("failed to fixup CoreFP: %w", err)
	}
	maps.Insert(symbolTable, maps.All(corefpO.SymbolTable()))

	if err := commercecoreO.Fixup(commercecoreAddr, symbolTable); err != nil {
		return nil, fmt.Errorf("failed to fixup CommerceCore: %w", err)
	}
	maps.Insert(symbolTable, maps.All(commercecoreO.SymbolTable()))

	if err := commercekitO.Fixup(commercekitAddr, symbolTable); err != nil {
		return nil, fmt.Errorf("failed to fixup CommerceKit: %w", err)
	}
	maps.Insert(symbolTable, maps.All(commercekitO.SymbolTable()))

	if err := storeagentO.Fixup(storeagentAddr, symbolTable); err != nil {
		return nil, fmt.Errorf("failed to fixup storeagent: %w", err)
	}
	maps.Insert(symbolTable, maps.All(storeagentO.SymbolTable()))

	uc, err := unicorn.NewUnicorn(unicorn.ARCH_X86, unicorn.MODE_64)
	if err != nil {
		return nil, err
	}

	e := Emulator{
		corefpicxsO:   corefpicxsO,
		corefpO:       corefpO,
		commercecoreO: commercecoreO,
		commercekitO:  commercekitO,
		storeagentO:   storeagentO,
		symbolTable:   symbolTable,
		unicorn:       uc,
	}

	if err := e.hookRDTSC(); err != nil {
		return nil, fmt.Errorf("failed to hook RDTSC: %w", err)
	}

	if err := e.loadInterposer(ip); err != nil {
		return nil, fmt.Errorf("failed to load Interposer: %w", err)
	}

	if err := e.loadCoreFPICXS(); err != nil {
		return nil, fmt.Errorf("failed to load CoreFP.icxs: %w", err)
	}

	if err := e.load(e.corefpO, corefpAddr); err != nil {
		return nil, fmt.Errorf("failed to load CoreFP: %w", err)
	}

	if err := e.load(e.commercecoreO, commercecoreAddr); err != nil {
		return nil, fmt.Errorf("failed to load CommerceCore: %w", err)
	}

	if err := e.load(e.commercekitO, commercekitAddr); err != nil {
		return nil, fmt.Errorf("failed to load CommerceKit: %w", err)
	}

	if err := e.load(e.storeagentO, storeagentAddr); err != nil {
		return nil, fmt.Errorf("failed to load storeagent: %w", err)
	}

	if err := e.setupCoreFPICXS(); err != nil {
		return nil, fmt.Errorf("failed to setup CoreFP.icxs: %w", err)
	}

	if err := e.setupCoreFP(); err != nil {
		return nil, fmt.Errorf("failed to setup CoreFP: %w", err)
	}

	if err := e.setupHeap(); err != nil {
		return nil, fmt.Errorf("failed to setup heap: %w", err)
	}

	if err := e.setupStack(); err != nil {
		return nil, fmt.Errorf("failed to setup stack: %w", err)
	}

	if err := e.setupStart(); err != nil {
		return nil, fmt.Errorf("failed to setup start: %w", err)
	}

	if err := e.setupTrace(); err != nil {
		return nil, fmt.Errorf("failed to setup trace: %w", err)
	}

	return &e, nil
}

func (e *Emulator) DumpHeap() ([]byte, error) {
	heapHead, err := e.unicorn.MemRead(e.symbolTable["_malloc.heap_head"], 8)
	if err != nil {
		return nil, err
	}
	heapHeadAddr := binary.LittleEndian.Uint64(heapHead)

	heapTail, err := e.unicorn.MemRead(e.symbolTable["_malloc.heap_tail"], 8)
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

func (e *Emulator) loadInterposer(ip *interposer.Interposer) error {
	pageSet := make(map[uint64]struct{})

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

			if _, ok := pageSet[addr]; ok {
				continue
			}
			pageSet[addr] = struct{}{}

			if err := e.unicorn.MemMap(addr, pageSize); err != nil {
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

func (e *Emulator) loadCoreFPICXS() error {
	size := e.corefpicxsO.Size()
	size += (pageSize - size) % pageSize

	if err := e.unicorn.MemMap(corefpicxsAddr, size); err != nil {
		return err
	}

	data := e.corefpicxsO.Data()
	return e.unicorn.MemWrite(corefpicxsAddr, data)
}

func (e *Emulator) load(o *library.Object, baseAddr uint64) error {
	virtSize, mappings := o.LoadInformation()

	virtSize += (pageSize - virtSize) % pageSize
	if err := e.unicorn.MemMap(baseAddr, virtSize); err != nil {
		return err
	}

	for addr, data := range mappings {
		addr += baseAddr
		if err := e.unicorn.MemWrite(addr, data); err != nil {
			return err
		}
	}

	return nil
}

func (e *Emulator) setupCoreFPICXS() error {
	var addr [8]byte
	binary.LittleEndian.PutUint64(addr[:], corefpicxsAddr)

	if err := e.unicorn.MemWrite(e.symbolTable["_read.CoreFP$ICXS_data"], addr[:]); err != nil {
		return err
	}

	var size [8]byte
	binary.LittleEndian.PutUint64(size[:], e.corefpicxsO.Size())

	if err := e.unicorn.MemWrite(e.symbolTable["_read.CoreFP$ICXS_size"], size[:]); err != nil {
		return err
	}

	return nil
}

func (e *Emulator) setupCoreFP() error {
	for _, sym := range []string{
		library.SymbolFairPlayUnknown0.String(),
		library.SymbolFairPlayUnknown1.String(),
		library.SymbolFairPlayUnknown2.String(),
		library.SymbolFairPlayUnknown3.String(),
		library.SymbolFairPlayUnknown4.String(),
		library.SymbolFairPlayUnknown5.String(),
	} {
		var addr [8]byte
		binary.LittleEndian.PutUint64(addr[:], e.symbolTable[sym])

		if err := e.unicorn.MemWrite(e.symbolTable["_dlsym."+sym], addr[:]); err != nil {
			return err
		}
	}

	return nil
}

func (e *Emulator) setupHeap() error {
	if err := e.unicorn.MemMap(heapHeadAddr, heapSize); err != nil {
		return err
	}

	var addrData [8]byte
	binary.LittleEndian.PutUint64(addrData[:], heapHeadAddr)

	var sizeData [8]byte
	binary.LittleEndian.PutUint64(sizeData[:], heapSize)

	if err := e.unicorn.MemWrite(e.symbolTable["_malloc.heap_head"], addrData[:]); err != nil {
		return err
	}

	if err := e.unicorn.MemWrite(e.symbolTable["_malloc.heap_tail"], addrData[:]); err != nil {
		return err
	}

	if err := e.unicorn.MemWrite(e.symbolTable["_malloc.heap_size"], sizeData[:]); err != nil {
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
		reverseTable := make(map[uint64]string, len(e.symbolTable))
		for sym, addr := range e.symbolTable {
			reverseTable[addr] = sym
		}

		hookCode := func(uc unicorn.Unicorn, addr uint64, size uint32) {
			if sym, ok := reverseTable[addr]; ok {
				_, _ = fmt.Fprintf(trace.Output, "->  0x%x <+?>: %s\n", addr, sym)
				return
			}

			if trace.Flags&trace.FlagOnCodeHookDisassemble != 0 {
				instruction, err := uc.MemRead(addr, uint64(size))
				if err != nil {
					logger.Print(err)
				}

				_, _ = fmt.Fprintf(trace.Output, "->  0x%x <+?>: ", addr)
				for _, b := range instruction {
					_, _ = fmt.Fprintf(trace.Output, "%02x ", b)
				}
				_, _ = fmt.Fprintf(trace.Output, "\n")
			}
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
