package emulator

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlaySAPTeardown(ctxRef *definitions.FPSAPContextOpaqueRef) error {
	var (
		arg0 = ctxRef.GetAddr()
	)
	if err := e.pushArgs(arg0); err != nil {
		return err
	}

	funcAddr := mescalAddr + e.mescalO.SymbolAddress(library.SymbolFairPlaySAPTeardown)
	if err := e.unicorn.RegWrite(unicorn.X86_REG_RAX, funcAddr); err != nil {
		return err
	}

	if err := e.unicorn.Start(startAddr, finishAddr); err != nil {
		return err
	}

	rax, err := e.unicorn.RegRead(unicorn.X86_REG_RAX)
	if err != nil {
		return err
	}

	if rax != 0 {
		return fmt.Errorf("FairPlaySAPTeardown: %d", int32(rax))
	}

	return nil
}
