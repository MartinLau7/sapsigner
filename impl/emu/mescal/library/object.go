package library

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/types"
)

type Object struct {
	data     []byte
	size     uint64
	baseAddr uint64
	segments []*macho.Segment
	symTable map[Symbol]uint64
	bindInfo []types.Bind
	rebaInfo []types.Rebase
}

func NewCoreFPICXSObject(data []byte) (*Object, error) {
	o := Object{
		data:     data,
		size:     uint64(len(data)),
		symTable: map[Symbol]uint64{},
	}

	return &o, nil
}

func NewCoreFPObject(data []byte) (*Object, error) {
	o := Object{
		symTable: map[Symbol]uint64{
			SymbolFairPlayUnknown0: 0,
			SymbolFairPlayUnknown1: 0,
			SymbolFairPlayUnknown2: 0,
			SymbolFairPlayUnknown3: 0,
			SymbolFairPlayUnknown4: 0,
			SymbolFairPlayUnknown5: 0,
		},
	}

	fatFile, err := macho.NewFatFile(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var file *macho.File
	for _, fatArch := range fatFile.Arches {
		if fatArch.CPU == types.CPUAmd64 {
			o.data = data[fatArch.Offset : fatArch.Offset+fatArch.Size]
			o.size = uint64(fatArch.Size)

			file = fatArch.File
			break
		}
	}

	o.baseAddr = file.GetBaseAddress()
	o.segments = file.Segments()

	for k := range o.symTable {
		name := k.String()

		addr, err := file.FindSymbolAddress(name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		addr -= o.baseAddr

		o.symTable[k] = addr
	}

	bindInfo, err := file.GetBindInfo()
	if err != nil {
		return nil, err
	}
	o.bindInfo = bindInfo

	rebaInfo, err := file.GetRebaseInfo()
	if err != nil {
		return nil, err
	}
	o.rebaInfo = rebaInfo

	return &o, nil
}

func NewCommerceCoreObject(data []byte) (*Object, error) {
	o := Object{
		symTable: map[Symbol]uint64{
			SymbolFairPlayGetMacAddress: 0,
		},
	}

	fatFile, err := macho.NewFatFile(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var file *macho.File
	for _, fatArch := range fatFile.Arches {
		if fatArch.CPU == types.CPUAmd64 {
			o.data = data[fatArch.Offset : fatArch.Offset+fatArch.Size]
			o.size = uint64(fatArch.Size)

			file = fatArch.File
			break
		}
	}

	o.baseAddr = file.GetBaseAddress()
	o.segments = file.Segments()

	for k := range o.symTable {
		name := k.String()

		addr, err := file.FindSymbolAddress(name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		addr -= o.baseAddr

		o.symTable[k] = addr
	}

	bindInfo, err := file.GetBindInfo()
	if err != nil {
		return nil, err
	}
	o.bindInfo = bindInfo

	rebaInfo, err := file.GetRebaseInfo()
	if err != nil {
		return nil, err
	}
	o.rebaInfo = rebaInfo

	return &o, nil
}

func NewCommerceKitObject(data []byte) (*Object, error) {
	o := Object{
		data: data,
		size: uint64(len(data)),
		symTable: map[Symbol]uint64{
			SymbolFairPlayDisposeStorage: 0,
			SymbolFairPlaySAPExchange:    0,
			SymbolFairPlaySAPInit:        0,
			SymbolFairPlaySAPPrime:       0,
			SymbolFairPlaySAPSign:        0,
			SymbolFairPlaySAPTeardown:    0,
			SymbolFairPlaySAPVerify:      0,
		},
	}

	file, err := macho.NewFile(bytes.NewReader(o.data))
	if err != nil {
		return nil, err
	}

	o.baseAddr = file.GetBaseAddress()
	o.segments = file.Segments()

	for k := range o.symTable {
		name := k.String()

		addr, err := file.FindSymbolAddress(name)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}
		addr -= o.baseAddr

		o.symTable[k] = addr
	}

	bindInfo, err := file.GetBindInfo()
	if err != nil {
		return nil, err
	}
	o.bindInfo = bindInfo

	rebaInfo, err := file.GetRebaseInfo()
	if err != nil {
		return nil, err
	}
	o.rebaInfo = rebaInfo

	return &o, nil
}

func NewStoreAgentObject(data []byte) (*Object, error) {
	o := Object{
		data: data,
		size: uint64(len(data)),
	}

	file, err := macho.NewFile(bytes.NewReader(o.data))
	if err != nil {
		return nil, err
	}

	o.baseAddr = file.GetBaseAddress()
	o.segments = file.Segments()

	o.symTable = map[Symbol]uint64{
		SymbolFairPlayGlobalContextInit:  0x1000C5FC0 - o.baseAddr,
		SymbolFairPlayKBSyncDataWithDSID: 0x1000C93C0 - o.baseAddr,
	}

	bindInfo, err := file.GetBindInfo()
	if err != nil {
		return nil, err
	}
	o.bindInfo = bindInfo

	rebaInfo, err := file.GetRebaseInfo()
	if err != nil {
		return nil, err
	}
	o.rebaInfo = rebaInfo

	return &o, nil
}

func (o *Object) Data() []byte {
	return o.data
}

func (o *Object) Size() uint64 {
	return o.size
}

func (o *Object) SymbolAddress(symbol Symbol) uint64 {
	return o.symTable[symbol]
}

func (o *Object) SymbolTable() map[string]uint64 {
	symTable := make(map[string]uint64, len(o.symTable))
	for sym, addr := range o.symTable {
		symTable[sym.String()] = addr
	}

	return symTable
}

func (o *Object) LoadInformation() (uint64, map[uint64][]byte) {
	virtSize := uint64(0)
	mappings := make(map[uint64][]byte, len(o.segments))

	var pageZero uint64
	for _, segment := range o.segments {
		if segment.Name == "__PAGEZERO" {
			pageZero = segment.Memsz
			continue
		}

		virtSize += segment.Memsz
		mappings[segment.Addr-pageZero] = o.data[segment.Offset : segment.Offset+segment.Filesz]
	}

	return virtSize, mappings
}

func (o *Object) Fixup(baseAddr uint64, symbolTable map[string]uint64) error {
	if err := o.rebasePointers(baseAddr); err != nil {
		return err
	}

	if err := o.bindReferences(symbolTable); err != nil {
		return err
	}

	o.baseAddr = baseAddr

	for sym := range o.symTable {
		o.symTable[sym] += baseAddr
	}

	return nil
}

func (o *Object) rebasePointers(baseAddr uint64) error {
	for _, rebase := range o.rebaInfo {
		if t := rebase.Type; t != types.REBASE_TYPE_POINTER {
			return fmt.Errorf("rebase type not supported: %d", t)
		}

		addr := rebase.Value
		addr -= o.baseAddr
		addr += baseAddr

		var addrData [8]byte
		binary.LittleEndian.PutUint64(addrData[:], addr)

		offs := rebase.Offset
		offs += rebase.Start
		offs -= o.baseAddr

		for _, b := range addrData {
			o.data[offs] = b

			offs++
		}
	}

	return nil
}

func (o *Object) bindReferences(symbolTable map[string]uint64) error {
	for _, bind := range o.bindInfo {
		addr, ok := symbolTable[bind.Name]
		if !ok {
			continue
		}
		addr += uint64(bind.Addend)

		if k, t := bind.Kind, bind.Type; k == types.BIND_KIND && t != types.BIND_TYPE_POINTER {
			return fmt.Errorf("bind type not supported: %d", t)
		}

		var addrData [8]byte
		binary.LittleEndian.PutUint64(addrData[:], addr)

		offs := bind.SegOffset
		offs += bind.SegStart

		for _, b := range addrData {
			o.data[offs] = b

			offs++
		}
	}

	return nil
}
