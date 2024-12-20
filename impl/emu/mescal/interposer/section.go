package interposer

type Section struct {
	addr uint64
	size uint64
	data []byte
}

func newSection(addr uint64, size uint64, data []byte) *Section {
	return &Section{
		addr: addr,
		size: size,
		data: data,
	}
}

func (s *Section) Addr() uint64 {
	return s.addr
}

func (s *Section) Size() uint64 {
	return s.size
}

func (s *Section) Data() []byte {
	return s.data
}
