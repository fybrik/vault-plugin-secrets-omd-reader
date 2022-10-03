// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omsecrets

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"

	omclient "fybrik.io/vault-plugin-secrets-omd-reader/openmetadata"
)

// backend wraps the backend framework
type secretsReaderBackend struct {
	*framework.Backend
	OMSecretReader OpenMetadataSecretsReader
}

var _ logical.Factory = Factory

// Factory configures and returns the plugin backends
func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b, err := newBackend()
	if err != nil {
		return nil, err
	}

	if conf == nil {
		return nil, fmt.Errorf("configuration passed into backend is nil")
	}

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return b, nil
}

func newBackend() (*secretsReaderBackend, error) { //nolint
	b := &secretsReaderBackend{
		OMSecretReader: OpenMetadataSecretsReader{client: &omclient.OMClient{}},
	}

	b.Backend = &framework.Backend{
		Help: strings.TrimSpace(backendHelp),
		// TypeLogical indicates that the backend (plugin) is a secret provider.
		BackendType: logical.TypeLogical,
		// Define the path for which this backend will respond.
		Paths: []*framework.Path{
			pathSecrets(b),
		},
	}

	return b, nil
}
