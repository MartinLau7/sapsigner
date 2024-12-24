module github.com/t0rr3sp3dr0/sapsigner/impl/emu

go 1.23

require (
	github.com/blacktop/go-apfs v1.0.18
	github.com/blacktop/go-macho v1.1.234
	github.com/unicorn-engine/unicorn v0.0.0-20241221030228-28990888443e
	howett.net/plist v1.0.1
	howett.net/ranger v0.0.0-20171016084633-e2e137620847
)

require (
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/apex/log v1.9.0 // indirect
	github.com/blacktop/go-dwarf v1.0.10 // indirect
	github.com/blacktop/go-plist v1.0.2 // indirect
	github.com/blacktop/lzfse-cgo v1.1.19 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/vbauerster/mpb/v7 v7.5.3 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace (
	github.com/blacktop/go-apfs v1.0.18 => github.com/t0rr3sp3dr0/go-apfs v1.0.19-0.20241224073602-4e50e3163b28
	github.com/blacktop/go-macho v1.1.234 => github.com/t0rr3sp3dr0/go-macho v0.0.0-20241224072836-20948aaf41de
)
