# Force Go Modules
GO111MODULE = on

GOCC ?= go
GOFLAGS ?=

# make reproducible
GOFLAGS += -asmflags=all=-trimpath="$(GOPATH)" -gcflags=all=-trimpath="$(GOPATH)"

.PHONY: install build

go.mod: FORCE
	./set-target.sh $(IPFS_VERSION)

FORCE:

supercluster: main.go
	$(GOCC) build $(GOFLAGS) -o "$@" "$<"
	chmod +x "$@"

build: supercluster
	@echo "Built $<"

install: build
	mkdir -p "$(IPFS_PATH)/plugins/"
	cp -f supercluster-plugin.so "$(IPFS_PATH)/plugins/supercluster-plugin.so"
