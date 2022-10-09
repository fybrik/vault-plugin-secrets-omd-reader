// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omsecrets

import (
	"context"
	"errors"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const secretsPrefix = "secret_name"

// pathSecrets returns the path configuration for reading openmetadata secrets.
func pathSecrets(b *secretsReaderBackend) *framework.Path {
	return &framework.Path{
		Pattern: framework.MatchAllRegex(secretsPrefix),

		Fields: map[string]*framework.FieldSchema{
			secretsPrefix: {
				Type:        framework.TypeString,
				Description: "Specifies the name of the openmetadata secret.",
				Query:       true,
				Required:    true,
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation: b.handleRead,
		},
		HelpDescription: pathInvalidHelp,
	}
}

const MissingSecretName string = "missing secret name"

// handleRead handles a read request: it extracts the secret name
// and returns the secret content if no error occured.
func (b *secretsReaderBackend) handleRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	secretName := data.Get(secretsPrefix).(string)
	b.Logger().Info("In handleRead() secretName: " + secretName)

	if secretName == "" {
		resp := logical.ErrorResponse(MissingSecretName)
		return resp, errors.New(MissingSecretName)
	}

	fetchedData, err := b.OMSecretReader.GetSecret(ctx, secretName, b.Logger())
	if err != nil {
		resp := logical.ErrorResponse("Error reading the secret data: " + err.Error())
		return resp, err
	}

	// Generate the response
	resp := &logical.Response{
		Data: fetchedData,
	}

	return resp, nil
}

var backendHelp string = `
This backend reads openmetadata secrets.`

var pathInvalidHelp string = backendHelp + `

## PATHS

The following paths are supported by this backend. To view help for
any of the paths below, use the help command with any route matching
the path pattern. Note that depending on the policy of your auth token,
you may or may not be able to access certain paths.

{{range .Paths}}{{indent 4 .Path}}
{{indent 8 .Help}}

{{end}}
`
