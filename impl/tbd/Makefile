override SHELL := $(SHELL) -o 'pipefail'

all: build
.PHONY: all

build: sapsigner.out
.PHONY: build

clean:
	rm -Rf *.out
.PHONY: clean

sapsigner.out: *.c *.h *.tbd
	clang -L. -O2 -Wall -Wextra -Wpedantic -lcurl -lmescal -lsasl2 -o $@ *.c
