package definitions

import (
	"encoding/binary"
	"fmt"
)

// FairPlaySAPExchangeVersion is a Go representation of the C type FairPlaySAPExchangeVersion_.
//
//	enum FairPlaySAPExchangeVersion_ {
//		FairPlaySAPExchangeVersion_Regular = 200,
//		FairPlaySAPExchangeVersion_Prime = 210,
//	};
//
//	typedef enum FairPlaySAPExchangeVersion_ FairPlaySAPExchangeVersion_;
type FairPlaySAPExchangeVersion [4]byte

var (
	FairPlaySAPExchangeVersionRegular = &FairPlaySAPExchangeVersion{0xC8, 0x00, 0x00, 0x00}
	FairPlaySAPExchangeVersionPrime   = &FairPlaySAPExchangeVersion{0xD2, 0x00, 0x00, 0x00}
)

func (fpsapev *FairPlaySAPExchangeVersion) Bytes() []byte {
	return fpsapev[:]
}

func (fpsapev *FairPlaySAPExchangeVersion) GetVal() uint32 {
	return binary.LittleEndian.Uint32(fpsapev[0:4])
}

func (fpsapev *FairPlaySAPExchangeVersion) SetVal(val uint32) {
	binary.LittleEndian.PutUint32(fpsapev[0:4], val)
}

var _ fmt.Stringer = (*FairPlaySAPExchangeVersion)(nil)

func (fpsapev *FairPlaySAPExchangeVersion) String() string {
	return fmt.Sprintf("%d", fpsapev.GetVal())
}
