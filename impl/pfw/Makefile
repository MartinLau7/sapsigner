override SHELL := $(SHELL) -o 'pipefail'

all: build
.PHONY: all

build: sapsigner.out
.PHONY: build

clean:
	rm -Rf *.out .build .swiftpm Package.resolved
.PHONY: clean

sapsigner.out: FORCE
	swift build -c release --arch arm64 --arch x86_64
	ln -Lf .build/apple/Products/Release/sapsigner $@
FORCE:
