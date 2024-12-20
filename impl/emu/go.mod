module github.com/t0rr3sp3dr0/sapsigner/impl/emu

go 1.23

require (
	github.com/blacktop/go-macho v1.1.233
	github.com/korylprince/go-cpio-odc v0.9.4
	github.com/unicorn-engine/unicorn v0.0.0-20240926111503-d568885d64c8
	howett.net/plist v1.0.1
)

require github.com/blacktop/go-dwarf v1.0.10 // indirect

replace (
	github.com/blacktop/go-macho v1.1.233 => github.com/t0rr3sp3dr0/go-macho v0.0.0-20241203062244-1a073aec0d7b
	github.com/korylprince/go-cpio-odc v0.9.4 => github.com/t0rr3sp3dr0/go-cpio-odc v0.9.5-0.20241204090443-a994b3aed6ee
	github.com/unicorn-engine/unicorn v0.0.0-20240926111503-d568885d64c8 => github.com/t0rr3sp3dr0/unicorn v0.0.0-20241220005206-b8204bd55907
)
