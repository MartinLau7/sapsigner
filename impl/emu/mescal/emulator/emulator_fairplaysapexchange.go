package emulator

import (
	"fmt"
	"math"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/definitions"
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/library"
)

func (e *Emulator) FairPlaySAPExchange(xVer *definitions.FairPlaySAPExchangeVersion, hwInfo *definitions.FairPlayHwInfo, ctxRef *definitions.FPSAPContextOpaqueRef, iBuf []byte) ([]byte, int32, error) {
	iLen := len(iBuf)
	if iLen > math.MaxUint32 {
		return nil, 0, fmt.Errorf("FairPlaySAPSign: input buffer too large")
	}

	hwInfoBridge, err := NewBridge(e.unicorn, hwInfo.Bytes())
	if err != nil {
		return nil, 0, err
	}
	defer hwInfoBridge.Close()

	iBufBridge, err := NewBridge(e.unicorn, iBuf)
	if err != nil {
		return nil, 0, err
	}
	defer iBufBridge.Close()

	var oBuf Buffer

	oPtrBridge, err := NewBridge(e.unicorn, oBuf.AddrBytes())
	if err != nil {
		return nil, 0, err
	}
	defer oPtrBridge.Close()

	oLenBridge, err := NewBridge(e.unicorn, oBuf.SizeBytes())
	if err != nil {
		return nil, 0, err
	}
	defer oLenBridge.Close()

	var rc Int32

	rcBridge, err := NewBridge(e.unicorn, rc.Bytes())
	if err != nil {
		return nil, 0, err
	}
	defer rcBridge.Close()

	var (
		arg0 = uint64(xVer.GetVal())
		arg1 = hwInfoBridge.Addr()
		arg2 = ctxRef.GetAddr()
		arg3 = iBufBridge.Addr()
		arg4 = uint64(iLen)
		arg5 = oPtrBridge.Addr()
		arg6 = oLenBridge.Addr()
		arg7 = rcBridge.Addr()
	)
	if err := e.pushArgs(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7); err != nil {
		return nil, 0, err
	}

	funcAddr := mescalAddr + e.mescalO.SymbolAddress(library.SymbolFairPlaySAPExchange)
	if err := e.unicorn.RegWrite(unicorn.X86_REG_RAX, funcAddr); err != nil {
		return nil, 0, err
	}

	if err := e.unicorn.Start(startAddr, finishAddr); err != nil {
		return nil, 0, err
	}

	rax, err := e.unicorn.RegRead(unicorn.X86_REG_RAX)
	if err != nil {
		return nil, 0, err
	}

	if rax != 0 {
		return nil, 0, fmt.Errorf("FairPlaySAPExchange: %d", int32(rax))
	}

	oDat, err := oBuf.Read(e.unicorn)
	if err != nil {
		return nil, 0, err
	}

	if err := e.FairPlayDisposeStorage(oPtrBridge.Addr()); err != nil {
		return nil, 0, err
	}

	rcDat := rc.GetVal()

	return oDat, rcDat, nil
}
