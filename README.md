# vault-plugin-secrets-omd-reader
Plugin for HashiCorp Vault which reads credential secrets from OpenMetadata

Requirements:

    make
    golang 1.19 and above
    Vault CLI utility (tested on Vault version v1.11.3)

## Quick Start

Begin by setting the `OM_SERVER_URL` environment variable to point to the OpenMetadata Rest API URL. For instance:
```
export OM_SERVER_URL="http://localhost:8585/api"
```

Next, build the plugin binary and start the Vault dev server:
```
make source-build
vault server -dev -dev-root-token-id=root -dev-plugin-dir=./vault/plugins
```

This Vault plugin will give you access to credentials of OpenMetadata database services.

Suppose your OpenMetadata deployment has a database service called `openmetadata-s3`. To get its credentials, open a new terminal window and run the following commands:

```
# Open a new terminal window and export Vault dev server http address
$ export VAULT_ADDR='http://127.0.0.1:8200'

# Enable the the plugin
$ vault secrets enable -path=om-secrets-reader vault-plugin-secrets-omd-reader

# Read the secret credentials of the openmetadata-s3 database service from OM:
$ vault read om-secrets-reader/openmetadata-s3
Key           Value
---           -----
access_key    myaccesskey
secret_key    mysecretkey
```
