COMMIT ?= $(shell git rev-parse --short HEAD)
VERSION ?= $(shell git describe --tags --dirty='+dev' --abbrev=0 || echo v0)
PERMALINK ?= $(shell git name-rev --name-only --tags --no-undefined HEAD &> /dev/null && echo latest || echo canary)

PKG = github.com/deislabs/porter-azure
LDFLAGS = -w -X $(PKG)/pkg.Version=$(VERSION) -X $(PKG)/pkg.Commit=$(COMMIT)
XBUILD = GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -ldflags '$(LDFLAGS)'

REGISTRY ?= $(USER)

build:
	mkdir -p bin/mixins/azure
	$(XBUILD) -o bin/mixins/azure/azure ./cmd/azure
	GOOS=linux $(XBUILD) -o bin/mixins/azure/azure-runtime ./cmd/azure

quickbuild:
	go build -o bin/mixins/azure/azure ./cmd/azure

test: clean test-unit

test-unit: build
	go test ./...

clean:
	-rm -fr bin/
