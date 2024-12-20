package emulator

import (
	"encoding/binary"
	"fmt"
)

type Int32 [4]byte

func (i32 *Int32) Bytes() []byte {
	return i32[:]
}

func (i32 *Int32) GetVal() int32 {
	return int32(binary.LittleEndian.Uint32(i32[0:4]))
}

func (i32 *Int32) SetVal(val int32) {
	binary.LittleEndian.PutUint32(i32[0:4], uint32(val))
}

var _ fmt.Stringer = (*Int32)(nil)

func (i32 *Int32) String() string {
	return fmt.Sprintf("%d", i32.GetVal())
}
