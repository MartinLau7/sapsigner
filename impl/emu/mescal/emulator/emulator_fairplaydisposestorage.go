package emulator

import (
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlayDisposeStorage(addr uint64) error {
	var (
		arg0 = addr
	)
	if err := e.pushArgs(arg0); err != nil {
		return err
	}

	funcAddr := commercekitAddr + e.commercekitO.SymbolAddress(library.SymbolFairPlayDisposeStorage)
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
		return fmt.Errorf("FairPlayDisposeStorage: %d", int32(rax))
	}

	return nil
}
