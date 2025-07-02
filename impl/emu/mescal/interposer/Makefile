override SHELL := $(SHELL) -o 'pipefail'

CONST_EXT := __TEXT,__const
CSTRING_EXT := __TEXT,__cstring
TEXT_EXT := __TEXT,__text
UNWINDINFO_EXT := __TEXT,__unwind_info

ALL_TARG := object/relt.go object/secb.go object/sech.go object/symt.go
SRC_TARG := $(wildcard replacements/*.c) main.c

INTERM_EXTS := $(CONST_EXT) $(CSTRING_EXT) $(TEXT_EXT) $(UNWINDINFO_EXT) o relt sech symt
INTERM_TARG := $(foreach EXT,$(INTERM_EXTS),$(patsubst %.c,%.$(EXT),$(SRC_TARG))) main.c
.INTERMEDIATE: $(INTERM_TARG)

all: $(ALL_TARG)
.PHONY: all

clean:
	rm -Rf $(patsubst %,'%',$(ALL_TARG)) $(patsubst %,'%',$(INTERM_TARG))
.PHONY: clean

test: $(patsubst %.c,%.o,$(SRC_TARG))
.PHONY: test

object/symt.go: object/symt.awk main.symt
	awk -f $(patsubst %,'%',$^) > '$@'

object/sech.go: object/sech.awk main.sech
	awk -f $(patsubst %,'%',$^) > '$@'

object/secb.go: object/secb.awk main.$(CONST_EXT) main.$(CSTRING_EXT) main.$(TEXT_EXT) main.$(UNWINDINFO_EXT)
	awk -f $(patsubst %,'%',$^) > '$@'

object/relt.go: object/relt.awk main.relt
	awk -f $(patsubst %,'%',$^) > '$@'

%.symt: %.o
	nm '$<' | grep -E '^[0-9a-z]+\s' > '$@'

%.sech: %.o
	objdump -h '$<' | grep -E '^\s*[0-9]+\s' > '$@'

%.relt: %.o
	objdump -r '$<' | grep -E '^[0-9a-z]+\s' > '$@'

%.$(CONST_EXT): %.o
	otool -X -s '__TEXT' '__const' '$<' > '$@'

%.$(CSTRING_EXT): %.o
	otool -X -s '__TEXT' '__cstring' '$<' > '$@'

%.$(TEXT_EXT): %.o
	otool -X -s '__TEXT' '__text' '$<' > '$@'

%.$(UNWINDINFO_EXT): %.o
	otool -X -s '__TEXT' '__unwind_info' '$<' > '$@'

%.o: %.c
	clang -O's' -W'all' -W'error' -W'extra' -W'pedantic' -arch 'x86_64' -c -f'freestanding' -f'no-PIC' -march='x86-64' -o'$@' -std='c2x' '$<'

%.c: %.h
	touch '$@'

main.c: $(filter-out main.c,$(SRC_TARG))
	xargs -n'1' 'printf' '#include "%s"\n' <<< '$^' > '$@'
