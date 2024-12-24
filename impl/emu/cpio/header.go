package cpio

import (
	"encoding/binary"
)

type Header struct {
	Magic    [6]byte
	Dev      [6]byte
	Ino      [6]byte
	Mode     [6]byte
	UID      [6]byte
	GID      [6]byte
	NLink    [6]byte
	RDev     [6]byte
	MTime    [11]byte
	NameSize [6]byte
	FileSize [11]byte
}

var (
	HeaderSize  = binary.Size(Header{})
	HeaderMagic = [6]byte{'0', '7', '0', '7', '0', '7'}
)

func IsValid(header Header, name []byte) bool {
	nameSize := len(name)
	return header.Magic == HeaderMagic && nameSize > 0 && name[nameSize-1] == 0
}
