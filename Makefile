override SHELL := $(SHELL) -o 'pipefail'

VARIANTS := $(notdir $(wildcard impl/*))

export

all: $(patsubst %,sapsigner-%.out,$(VARIANTS))
.PHONY: all

all(%):
	$(MAKE) -C impl/$% all
.PHONY: all(%)

build: $(patsubst %,build(%),$(VARIANTS))
.PHONY: build

build(%):
	$(MAKE) -C impl/$% build
.PHONY: build(%)

clean: $(patsubst %,clean(%),$(VARIANTS))
	rm -Rf *.out
.PHONY: clean

clean(%):
	$(MAKE) -C impl/$% clean
.PHONY: clean(%)

docker:
	docker build -t t0rr3sp3dr0/sapsigner .
.PHONY: docker

test: test(hack/docker-run.sh) $(patsubst %,test(impl/%/sapsigner.out),$(VARIANTS))
.PHONY: test

test(%): %
	$(eval override F := $(abspath $<))
	./hack/test-authenticate.sh '$(F)'
	./hack/test-signupWizard.sh '$(F)'
.PHONY: test(%)

impl/%/sapsigner.out: all(%)
	@ :

sapsigner-%.out: impl/%/sapsigner.out
	ln -Lf $< $@
