// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omclient

import (
	"context"
	"encoding/json"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"

	"fybrik.io/vault-plugin-secrets-omd-reader/pkg/utils"
)

const FybrikAccessKeyString = "access_key"
const FybrikSecretKeyString = "secret_key"

type OMClient struct {
}

// Return WKClient object
func NewOMClient() *client.APIClient {
	conf := client.Configuration{Servers: client.ServerConfigurations{
		client.ServerConfiguration{
			URL:         utils.GetOMServerURL(),
			Description: "Endpoint URL",
		},
	},
	}

	return client.NewAPIClient(&conf)
}

type Config struct {
	ConfigSource struct {
		SecurityConfig struct {
			AccessKey string `json:"awsAccessKeyId"`
			SecretKey string `json:"awsSecretAccessKey"`
		} `json:"securityConfig"`
	} `json:"configSource"`
}

func (o *OMClient) ExtractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error) {
	var c Config
	configBytes, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configBytes, &c)
	if err != nil {
		return nil, err
	}
	securityConfig := c.ConfigSource.SecurityConfig
	return map[string]interface{}{
		FybrikAccessKeyString: securityConfig.AccessKey,
		FybrikSecretKeyString: securityConfig.SecretKey,
	}, nil
}

func (o *OMClient) GetConnectionInformation(ctx context.Context, connectionName string) (*client.DatabaseService, error) {
	c := NewOMClient()
	databaseService, _, err := c.DatabaseServiceApi.GetDatabaseServiceByFQN(ctx, connectionName).Execute()
	if err != nil {
		return nil, err
	}
	return databaseService, nil
}
