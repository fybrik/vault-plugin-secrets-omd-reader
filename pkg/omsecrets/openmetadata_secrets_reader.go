// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omsecrets

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"

	omclient "fybrik.io/vault-plugin-secrets-omd-reader/pkg/openmetadata"
	dbtypes "fybrik.io/vault-plugin-secrets-omd-reader/pkg/openmetadata/database-types"
)

type OpenMetadataSecretsReader struct {
	client *omclient.OMClient
}

// GetSecret returns the content of openmetadata secret.
func (s *OpenMetadataSecretsReader) GetSecret(ctx context.Context, secretName string, log hclog.Logger) (map[string]interface{}, error) {
	nameToDatabaseStruct := map[string]dbtypes.Databasetype{
		"Datalake": dbtypes.S3{},
	}

	databaseService, err := s.client.GetConnectionInformation(ctx, secretName, log)
	if err != nil {
		return nil, err
	}
	serviceType := databaseService.GetServiceType()
	dt, found := nameToDatabaseStruct[serviceType]
	if !found {
		return nil, fmt.Errorf("Service type %s not recognized", serviceType)
	}
	config := databaseService.Connection.GetConfig()
	return dt.ExtractSecretsFromConfig(config)
}
