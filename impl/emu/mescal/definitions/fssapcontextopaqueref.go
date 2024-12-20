package definitions

import (
	"encoding/binary"
	"fmt"
)

// FPSAPContextOpaqueRef is a Go representation of the C type FPSAPContextOpaque_Ref.
//
//	struct FPSAPContextOpaque_;
//
//	typedef struct FPSAPContextOpaque_ FPSAPContextOpaque_;
//
//	typedef FPSAPContextOpaque_ *FPSAPContextOpaque_Ref;
type FPSAPContextOpaqueRef [8]byte

func (fpsapcor *FPSAPContextOpaqueRef) Bytes() []byte {
	return fpsapcor[:]
}

func (fpsapcor *FPSAPContextOpaqueRef) GetAddr() uint64 {
	return binary.LittleEndian.Uint64(fpsapcor[0:8])
}

func (fpsapcor *FPSAPContextOpaqueRef) SetAddr(addr uint64) {
	binary.LittleEndian.PutUint64(fpsapcor[0:8], addr)
}

var _ fmt.Stringer = (*FPSAPContextOpaqueRef)(nil)

func (fpsapcor *FPSAPContextOpaqueRef) String() string {
	return fmt.Sprintf("0x%016X", fpsapcor.GetAddr())
}
