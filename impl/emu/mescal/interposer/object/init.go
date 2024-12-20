//
//  init.go
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-30.
//

package object

import (
	"encoding/binary"
	"fmt"
)

func init() {
	if err := resolveRelocations(); err != nil {
		panic(err)
	}
}

func resolveRelocations() error {
	for k, v := range RelocationTable {
		addr := k
		typ_ := v[0]
		name := v[1]

		switch typ_ {
		case "X86_64_RELOC_BRANCH", "X86_64_RELOC_SIGNED":
			break
		default:
			return fmt.Errorf("relocation type not supported: %s", typ_)
		}

		sect, err := getSectionForAddress(addr)
		if err != nil {
			return err
		}

		offs := addr
		offs -= SectionHeaders[sect][0]

		relo := SymbolTable[name]
		relo -= k + 4

		var reloData [4]byte
		binary.LittleEndian.PutUint32(reloData[:], uint32(relo))

		data := SectionBytes[sect]
		for i, b := range reloData {
			data[offs+uint64(i)] = b
		}
	}

	return nil
}

func getSectionForAddress(addr uint64) (string, error) {
	for k, v := range SectionHeaders {
		name := k
		offs := v[0]
		size := v[1]

		if addr >= offs && addr < offs+size {
			return name, nil
		}
	}

	return "", fmt.Errorf("could not find section for address: 0x%016X", addr)
}
