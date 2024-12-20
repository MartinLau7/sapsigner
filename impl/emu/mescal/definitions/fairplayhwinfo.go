package definitions

import (
	"encoding/binary"
	"fmt"
)

// FairPlayHwInfo is a Go representation of the C type FairPlayHWInfo_.
//
//	struct FairPlayHWInfo_ {
//		unsigned int IDLength;
//		unsigned char ID[20];
//	};
//
//	typedef struct FairPlayHWInfo_ FairPlayHWInfo_;
type FairPlayHwInfo [24]byte

func (fphw *FairPlayHwInfo) Bytes() []byte {
	return fphw[:]
}

func (fphw *FairPlayHwInfo) GetId() []byte {
	id := fphw.getId()
	idLength := fphw.getIdLength()

	return id[:idLength]
}

func (fphw *FairPlayHwInfo) SetId(id []byte) {
	idLength := len(id)
	if idLength > 20 {
		panic("len(id) > 20")
	}

	fphw.setId(id)
	fphw.setIdLength(idLength)
}

func (fphw *FairPlayHwInfo) getId() [20]byte {
	return [20]byte(fphw[4:24])
}

func (fphw *FairPlayHwInfo) setId(id []byte) {
	copy(fphw[4:24], id[:])
}

func (fphw *FairPlayHwInfo) getIdLength() uint32 {
	return binary.LittleEndian.Uint32(fphw[0:4])
}

func (fphw *FairPlayHwInfo) setIdLength(idLength int) {
	binary.LittleEndian.PutUint32(fphw[0:4], uint32(idLength))
}

var _ fmt.Stringer = (*FairPlayHwInfo)(nil)

func (fphw *FairPlayHwInfo) String() string {
	return fmt.Sprintf("{IDLength:%d ID:[% X]}", fphw.getIdLength(), fphw.getId())
}
