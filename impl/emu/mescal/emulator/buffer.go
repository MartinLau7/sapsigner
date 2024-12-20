package emulator

import (
	"encoding/binary"
	"fmt"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type Buffer struct {
	addr [8]byte
	size [8]byte
}

func (b *Buffer) AddrBytes() []byte {
	return b.addr[:]
}

func (b *Buffer) SizeBytes() []byte {
	return b.size[:]
}

func (b *Buffer) Read(uc unicorn.Unicorn) ([]byte, error) {
	return uc.MemRead(b.getAddr(), b.getSize())
}

func (b *Buffer) getAddr() uint64 {
	return binary.LittleEndian.Uint64(b.addr[0:8])
}

func (b *Buffer) getSize() uint64 {
	return binary.LittleEndian.Uint64(b.size[0:8])
}

var _ fmt.Stringer = (*Buffer)(nil)

func (b *Buffer) String() string {
	return fmt.Sprintf("[%d]byte(0x%016X)", b.getAddr(), b.getSize())
}
