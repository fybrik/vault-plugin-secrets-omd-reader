GOARCH = amd64
OS = linux

.DEFAULT_GOAL := all

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_TAG ?= 0.0.0
DOCKER_NAME ?= vault-plugin-secrets-omd-reader

IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}

all: source-build

.PHONY: source-build
source-build:
	CGO_ENABLED=0 GOOS="$(OS)" GOARCH="$(GOARCH)" go build -o vault/plugins/vault-plugin-secrets-omd-reader pkg/cmd/vault-plugin-secrets-omd-reader/main.go

.PHONY: docker-build
docker-build: source-build
	docker build -f Dockerfile . -t ${IMG}

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: enable
enable:
	vault secrets enable -path=omd-secrets-reader vault-plugin-secrets-omd-reader 

.PHONY: clean
clean:
	rm -f ./vault/plugins/vault-plugin-secrets-omd-reader
