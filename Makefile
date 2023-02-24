# Force Go Modules
GO111MODULE = on

GOCC ?= go
GOFLAGS ?=
HOME_VAR = $(shell echo $$HOME)
SUPERCLUSTER_DIR = $(HOME)"/.supercluster"

# make reproducible
GOFLAGS += -asmflags=all=-trimpath="$(GOPATH)" -gcflags=all=-trimpath="$(GOPATH)"

.PHONY: install build

supercluster: main.go
ifeq ($(wildcard ./build),)
	./scripts/setup.sh
endif
	$(GOCC) build $(GOFLAGS) -o "./build/$@" "$<"

build: supercluster
	@echo "Built $<"

install: build
	@sudo cp build/supercluster /usr/local/bin/

ifeq "$(wildcard $(SUPERCLUSTER_DIR)/kubo)" ""
	@cp -r build/kubo ~/.supercluster
endif
