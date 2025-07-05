package definitions

import (
	"encoding/binary"
	"fmt"
)

// FairPlayGlobalContext is a Go representation of the C type FairPlayGlobalContext_.
//
//	typedef unsigned int FairPlayGlobalContext_;
type FairPlayGlobalContext [4]byte

func (fpgc *FairPlayGlobalContext) Bytes() []byte {
	return fpgc[:]
}

func (fpgc *FairPlayGlobalContext) GetVal() uint32 {
	return binary.LittleEndian.Uint32(fpgc[0:4])
}

func (fpgc *FairPlayGlobalContext) SetVal(val uint32) {
	binary.LittleEndian.PutUint32(fpgc[0:4], val)
}

var _ fmt.Stringer = (*FairPlayGlobalContext)(nil)

func (fpgc *FairPlayGlobalContext) String() string {
	return fmt.Sprintf("%d", fpgc.GetVal())
}
