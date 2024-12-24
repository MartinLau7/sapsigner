package cpio

import (
	"slices"
)

var (
	TrailerName   = [13]byte{'T', 'R', 'A', 'I', 'L', 'E', 'R', '!', '!', '!', '\x00', '\x00', '\x00'}
	TrailerHeader = Header{
		Magic:    HeaderMagic,
		Dev:      [6]byte{'0', '0', '0', '0', '0', '0'},
		Ino:      [6]byte{'0', '0', '0', '0', '0', '0'},
		Mode:     [6]byte{'0', '0', '0', '0', '0', '0'},
		UID:      [6]byte{'0', '0', '0', '0', '0', '0'},
		GID:      [6]byte{'0', '0', '0', '0', '0', '0'},
		NLink:    [6]byte{'0', '0', '0', '0', '0', '1'},
		RDev:     [6]byte{'0', '0', '0', '0', '0', '0'},
		MTime:    [11]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
		NameSize: [6]byte{'0', '0', '0', '0', '1', '3'},
		FileSize: [11]byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
	}
)

func IsTrailer(header Header, name []byte) bool {
	return header == TrailerHeader && slices.Equal(name, TrailerName[:])
}
