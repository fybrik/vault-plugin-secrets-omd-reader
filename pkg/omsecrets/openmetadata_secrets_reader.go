// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omsecrets

import (
	"context"

	"github.com/hashicorp/go-hclog"

	omclient "fybrik.io/vault-plugin-secrets-omd-reader/pkg/openmetadata"
)

type OpenMetadataSecretsReader struct {
	client *omclient.OMClient
}

// GetSecret returns the content of openmetadata secret.
func (s *OpenMetadataSecretsReader) GetSecret(ctx context.Context, secretName string, log hclog.Logger) (map[string]interface{}, error) {
	databaseService, err := s.client.GetConnectionInformation(ctx, secretName, log)
	if err != nil {
		return nil, err
	}
	config := databaseService.Connection.GetConfig()
	return s.client.ExtractSecretsFromConfig(config)
}
