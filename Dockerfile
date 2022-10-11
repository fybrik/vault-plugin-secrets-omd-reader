FROM alpine:latest

WORKDIR /
COPY vault/plugins/vault-plugin-secrets-omd-reader .
USER 10001
