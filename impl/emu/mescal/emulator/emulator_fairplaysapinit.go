package emulator

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlaySAPInit(hwInfo *definitions.FairPlayHwInfo) (*definitions.FPSAPContextOpaqueRef, error) {
	var ctxRef definitions.FPSAPContextOpaqueRef

	ctxRefBridge, err := NewBridge(e.unicorn, ctxRef.Bytes())
	if err != nil {
		return nil, err
	}
	defer ctxRefBridge.Close()

	hwInfoBridge, err := NewBridge(e.unicorn, hwInfo.Bytes())
	if err != nil {
		return nil, err
	}
	defer hwInfoBridge.Close()

	var (
		arg0 = ctxRefBridge.Addr()
		arg1 = hwInfoBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1); err != nil {
		return nil, err
	}

	funcAddr := e.commercekitO.SymbolAddress(library.SymbolFairPlaySAPInit)
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
		return nil, fmt.Errorf("FairPlaySAPInit: %d", int32(rax))
	}

	return &ctxRef, nil
}
