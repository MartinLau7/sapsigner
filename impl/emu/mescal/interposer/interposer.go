package interposer

import (
	"github.com/t0rr3sp3dr0/sapsigner/impl/emu/mescal/interposer/object"
)

type Interposer struct {
	sections []*Section
	symTable map[string]uint64
}

func NewInterposer(baseAddr uint64) *Interposer {
	var i Interposer

	symTable := make(map[string]uint64, len(object.SymbolTable))
	for k, v := range object.SymbolTable {
		symTable[k] = v + baseAddr
	}
	i.symTable = symTable

	sections := make([]*Section, 0, len(object.SectionHeaders))
	for k, v := range object.SectionHeaders {
		addr := v[0] + baseAddr
		size := v[1]
		data := object.SectionBytes[k]

		section := newSection(addr, size, data)
		sections = append(sections, section)
	}
	i.sections = sections

	return &i
}

func (i *Interposer) Sections() []*Section {
	return i.sections
}

func (i *Interposer) SymbolTable() map[string]uint64 {
	return i.symTable
}
