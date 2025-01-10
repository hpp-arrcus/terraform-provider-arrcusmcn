TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=arrcus.com
NAMESPACE=arrcus
NAME=arrcusmcn
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_arm64

default: install

build:
	go build -o ${BINARY}

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

local-install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

.PHONY: build local-install
