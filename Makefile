SHELL := $(SHELL) -o 'pipefail'

VARIANTS := $(notdir $(wildcard impl/*))

all: build
.PHONY: all

build: $(patsubst %,sapsigner-%.out,$(VARIANTS))
.PHONY: build

clean: $(patsubst %,clean(%),$(VARIANTS))
	rm -Rf *.out
.PHONY: clean

clean(%):
	$(MAKE) -C impl/$% SHELL='$(SHELL)' clean
.PHONY: clean(%)

docker:
	docker build -t t0rr3sp3dr0/sapsigner .
.PHONY: docker

test: test(hack/docker-run.sh) $(patsubst %,test(impl/%/sapsigner.out),$(VARIANTS))
.PHONY: test

test(%): %
	$(eval F := $(abspath $<))
	./hack/test-authenticate.sh '$(F)'
	./hack/test-signupWizard.sh '$(F)'
.PHONY: test(%)

impl/%/sapsigner.out:
	$(MAKE) -C impl/$(patsubst impl/%/sapsigner.out,%,$@) SHELL='$(SHELL)' sapsigner.out

sapsigner-%.out: impl/%/sapsigner.out
	ln -Lf $< $@
