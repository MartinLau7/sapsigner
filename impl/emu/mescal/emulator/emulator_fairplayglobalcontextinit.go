package emulator

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlayGlobalContextInit(hwInfo *definitions.FairPlayHwInfo) (*definitions.FairPlayGlobalContext, error) {
	hwInfoBridge, err := NewBridge(e.unicorn, hwInfo.Bytes())
	if err != nil {
		return nil, err
	}
	defer hwInfoBridge.Close()

	scInfo := []byte("/Users/Shared/SC Info\x00")

	scInfoBridge, err := NewBridge(e.unicorn, scInfo)
	if err != nil {
		return nil, err
	}
	defer scInfoBridge.Close()

	var globalContext definitions.FairPlayGlobalContext

	globalContextBridge, err := NewBridge(e.unicorn, globalContext.Bytes())
	if err != nil {
		return nil, err
	}
	defer globalContextBridge.Close()

	var (
		arg0 = uint64(0)
		arg1 = hwInfoBridge.Addr()
		arg2 = scInfoBridge.Addr()
		arg3 = globalContextBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1, arg2, arg3); err != nil {
		return nil, err
	}

	funcAddr := e.storeagentO.SymbolAddress(library.SymbolFairPlayGlobalContextInit)
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
		// Error -42076 -> looks to indicate problems with memory data
		// Error -42408 -> looks to indicate problems with CoreFP.icxs

		return nil, fmt.Errorf("FairPlayGlobalContextInit: %d", int32(rax))
	}

	return &globalContext, nil
}
