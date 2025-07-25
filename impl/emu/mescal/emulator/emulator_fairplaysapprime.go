package emulator

import (
	"fmt"
	"math"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlaySAPPrime(ctxRef *definitions.FPSAPContextOpaqueRef, pVer *definitions.FairPlaySAPPrimeVersion, iBuf []byte) ([]byte, error) {
	iLen := len(iBuf)
	if iLen > math.MaxUint32 {
		return nil, fmt.Errorf("FairPlaySAPPrime: input buffer too large")
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
		arg1 = uint64(pVer.GetVal())
		arg2 = iBufBridge.Addr()
		arg3 = uint64(iLen)
		arg4 = oPtrBridge.Addr()
		arg5 = oLenBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1, arg2, arg3, arg4, arg5); err != nil {
		return nil, err
	}

	funcAddr := e.commercekitO.SymbolAddress(library.SymbolFairPlaySAPPrime)
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
		return nil, fmt.Errorf("FairPlaySAPPrime: %d", int32(rax))
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
