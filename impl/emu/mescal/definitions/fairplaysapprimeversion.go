package definitions

import (
	"encoding/binary"
	"fmt"
)

// FairPlaySAPPrimeVersion is a Go representation of the C type FairPlaySAPPrimeVersion_.
//
//	enum FairPlaySAPPrimeVersion_ {
//		FairPlaySAPPrimeVersion_Regular = 200,
//		FairPlaySAPPrimeVersion_Prime = 210,
//	};
//
//	typedef enum FairPlaySAPPrimeVersion_ FairPlaySAPPrimeVersion_;
type FairPlaySAPPrimeVersion [4]byte

var (
	FairPlaySAPPrimeVersionRegular = &FairPlaySAPPrimeVersion{0x64, 0x00, 0x00, 0x00}
)

func (fpsapev *FairPlaySAPPrimeVersion) Bytes() []byte {
	return fpsapev[:]
}

func (fpsapev *FairPlaySAPPrimeVersion) GetVal() uint32 {
	return binary.LittleEndian.Uint32(fpsapev[0:4])
}

func (fpsapev *FairPlaySAPPrimeVersion) SetVal(val uint32) {
	binary.LittleEndian.PutUint32(fpsapev[0:4], val)
}

var _ fmt.Stringer = (*FairPlaySAPPrimeVersion)(nil)

func (fpsapev *FairPlaySAPPrimeVersion) String() string {
	return fmt.Sprintf("%d", fpsapev.GetVal())
}
