override SHELL := $(SHELL) -o 'pipefail'

UNAME_MACHINE := $(shell uname -m)
UNAME_SYSNAME := $(shell uname -s)

ifeq ($(UNAME_SYSNAME),Darwin)
	ifeq ($(UNAME_MACHINE),arm64)
		UNICORN_BLOB := sha256:8cb9b6e44b8f70a0e9cf1da0ba23dc7f9a0cbafaa5e2044f83166918dc5560e0
	else ifeq ($(UNAME_MACHINE),x86_64)
		UNICORN_BLOB := sha256:aeb5bf15f1613f87677b7bb539cac1bc9b3fcd339f06881437058ed19ac5313a
	endif
else ifeq ($(UNAME_SYSNAME),Linux)
	ifeq ($(UNAME_MACHINE),x86_64)
		UNICORN_BLOB := sha256:8935edc693fe1e8f5756c96913bf9de63aed9d22be914ef6a7b40dff96686e61
	endif
endif

UNICORN_VERSION := 2.1.1

ifneq ($(BUILD_ARGS),)
	override NULL :=
	override BUILD_ARGS := $(NULL) $(BUILD_ARGS)
endif

ifneq ($(BUILD_STATIC),)
	override BUILD_ARGS := $(BUILD_ARGS) -tags=static
	ifneq ($(UNAME_SYSNAME),Darwin)
		override BUILD_ARGS := $(BUILD_ARGS) -ldflags='-extldflags=-static -linkmode=external'
	endif
endif

all: vendor build
.PHONY: all

build: sapsigner.out kbsync.out
.PHONY: build

clean:
	rm -Rfv *.out include lib
.PHONY: clean

format:
	go fmt ./...
	go run github.com/daixiang0/gci write --custom-order --skip-generated --skip-vendor -s 'standard' -s 'prefix(golang.org/x)' -s 'blank' -s 'dot' -s 'default' -s 'localmodule' .
.PHONY: format

vendor: include/unicorn/unicorn.h lib/libunicorn.a
.PHONY: clean

include/unicorn/unicorn.h lib/libunicorn.a lib/libunicorn.dylib:
ifeq ($(UNICORN_BLOB),)
	$(warning Unable to vendor Unicorn $(UNICORN_VERSION) for $(UNAME_SYSNAME) $(UNAME_MACHINE))
else
	curl -H 'Accept: application/vnd.oci.image.index.v1+json' -H 'Authorization: Bearer QQ==' -Lf 'https://ghcr.io/v2/homebrew/core/unicorn/blobs/$(UNICORN_BLOB)' | zcat -f | tar --strip-components 2 -vx unicorn/$(UNICORN_VERSION)/include unicorn/$(UNICORN_VERSION)/lib
ifeq ($(UNAME_SYSNAME),Darwin)
	install_name_tool -id '@rpath/libunicorn.2.dylib' lib/libunicorn.2.dylib
endif
endif

%.out: FORCE
	CGO_LDFLAGS_ALLOW='-Wl,-rpath,[^,]+' go build$(BUILD_ARGS) -o='$@' -trimpath ./cmd/$(@:%.out=%)/...
FORCE:
