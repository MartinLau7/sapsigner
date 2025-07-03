package emulator

import (
	"fmt"
	"math"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlaySAPVerify(ctxRef *definitions.FPSAPContextOpaqueRef, iBuf []byte) ([]byte, error) {
	iLen := len(iBuf)
	if iLen > math.MaxUint32 {
		return nil, fmt.Errorf("FairPlaySAPVerify: input buffer too large")
	}

	iBufBridge, err := NewBridge(e.unicorn, iBuf)
	if err != nil {
		return nil, err
	}
	defer iBufBridge.Close()

	var oBuf Buffer

	oPtrBridge, err := NewBridge(e.unicorn, oBuf.AddrBytes())
	if err != nil {
		return nil, err
	}
	defer oPtrBridge.Close()

	oLenBridge, err := NewBridge(e.unicorn, oBuf.SizeBytes())
	if err != nil {
		return nil, err
	}
	defer oLenBridge.Close()

	var (
		arg0 = ctxRef.GetAddr()
		arg1 = iBufBridge.Addr()
		arg2 = uint64(iLen)
		arg3 = oPtrBridge.Addr()
		arg4 = oLenBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1, arg2, arg3, arg4); err != nil {
		return nil, err
	}

	funcAddr := e.commercekitO.SymbolAddress(library.SymbolFairPlaySAPVerify)
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
		return nil, fmt.Errorf("FairPlaySAPVerify: %d", int32(rax))
	}

	oDat, err := oBuf.Read(e.unicorn)
	if err != nil {
		return nil, err
	}

	if err := e.FairPlayDisposeStorage(oPtrBridge.Addr()); err != nil {
		return nil, err
	}

	return oDat, nil
}
