GOARCH = amd64
OS = linux

.DEFAULT_GOAL := all

all: source-build

.PHONY: source-build
source-build:
	CGO_ENABLED=0 GOOS="$(OS)" GOARCH="$(GOARCH)" go build -o vault/plugins/vault-plugin-secrets-omd-reader cmd/vault-plugin-secrets-omd-reader/main.go

.PHONY: enable
enable:
	vault secrets enable -path=omd-secrets-reader vault-plugin-secrets-omd-reader 

.PHONY: clean
clean:
	rm -f ./vault/plugins/vault-plugin-secrets-omd-reader
