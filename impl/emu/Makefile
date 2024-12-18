SHELL := $(SHELL) -o 'pipefail'

all: vendor build
.PHONY: all

build: sapsigner.out
.PHONY: build

clean:
	rm -Rfv *.out include lib
.PHONY: clean

vendor: include/unicorn/unicorn.h lib/libunicorn.a
.PHONY: clean

UNAME_MACHINE ?= $(shell uname -m)
UNAME_SYSNAME ?= $(shell uname -s)
ifeq ($(UNAME_SYSNAME),Linux)
	RPATH_DOT := \$$ORIGIN

	ifeq ($(UNAME_MACHINE),x86_64)
		UNICORN_BLOB := sha256:f45a8a230a4a63c46fc1d9a63e4d3a8ea33e6448051200bb202c23081efc96f2
	endif
else ifeq ($(UNAME_SYSNAME),Darwin)
	RPATH_DOT := @loader_path

	ifeq ($(UNAME_MACHINE),x86_64)
		UNICORN_BLOB := sha256:e092a3f6f7fbaa3483b3fa2b892da569adb0e3773aed3135aafe3eb9de132210
	else ifeq ($(UNAME_MACHINE),arm64)
		UNICORN_BLOB := sha256:19357aa9fc753df6791bf5a1764dbb2fdfeed8cdcc3a6b5c8174a052558cd8e2
	endif
endif

include/unicorn/unicorn.h lib/libunicorn.a:
	curl -H 'Accept: application/vnd.oci.image.index.v1+json' -H 'Authorization: Bearer QQ==' -Lf 'https://ghcr.io/v2/homebrew/core/unicorn/blobs/$(UNICORN_BLOB)' | zcat -f | tar --strip-components 2 -vx unicorn/2.0.1.post1/include unicorn/2.0.1.post1/lib
ifeq ($(UNAME_SYSNAME),Darwin)
	install_name_tool -id @rpath/libunicorn.2.dylib lib/libunicorn.2.dylib
endif

sapsigner.out: *.c *.h
	clang -I include -I /opt/homebrew/include -I /usr/local/include -L lib -L /opt/homebrew/lib -L /usr/local/lib -O2 -Wall -Wextra -Wpedantic -lcurl -lsasl2 -lunicorn -rpath $(RPATH_DOT) -rpath $(RPATH_DOT)/lib -rpath /opt/homebrew/lib -rpath /usr/local/lib -o $@ *.c
