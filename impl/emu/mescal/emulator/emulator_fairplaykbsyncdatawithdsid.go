package emulator

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlayKBSyncDataWithDSID(globalContext *definitions.FairPlayGlobalContext, dsid uint64) ([]byte, error) {
	var buf Buffer

	ptrBridge, err := NewBridge(e.unicorn, buf.AddrBytes())
	if err != nil {
		return nil, err
	}
	defer ptrBridge.Close()

	lenBridge, err := NewBridge(e.unicorn, buf.SizeBytes())
	if err != nil {
		return nil, err
	}
	defer lenBridge.Close()

	var (
		arg0 = uint64(globalContext.GetVal())
		arg1 = dsid
		arg2 = uint64(0)
		arg3 = uint64(1)
		arg4 = ptrBridge.Addr()
		arg5 = lenBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1, arg2, arg3, arg4, arg5); err != nil {
		return nil, err
	}

	funcAddr := e.storeagentO.SymbolAddress(library.SymbolFairPlayKBSyncDataWithDSID)
	if err := e.unicorn.RegWrite(unicorn.X86_REG_RAX, funcAddr); err != nil {
		return nil, err
	}

	if err := e.unicorn.Start(startAddr, finishAddr); err != nil {
		return nil, err
	}

	rax, err := e.unicorn.RegRead(unicorn.X86_REG_RAX)
	if err != nil {
		return nil, err
	}

	if rax != 0 {
		return nil, fmt.Errorf("FairPlayKBSyncDataWithDSID: %d", int32(rax))
	}

	dat, err := buf.Read(e.unicorn)
	if err != nil {
		return nil, err
	}

	if err := e.FairPlayDisposeStorage(ptrBridge.Addr()); err != nil {
		return nil, err
	}

	return dat, nil
}
