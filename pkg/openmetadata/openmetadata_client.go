// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omclient

import (
	"context"
	"fmt"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	"github.com/hashicorp/go-hclog"

	"fybrik.io/vault-plugin-secrets-omd-reader/pkg/utils"
)

type OMClient struct {
}

// Return WKClient object
func NewOMClient(ctx context.Context, logger hclog.Logger) (*client.APIClient, error) {
	url, username, password := utils.GetEnvironmentVariables()
	conf := client.Configuration{Servers: client.ServerConfigurations{
		client.ServerConfiguration{
			URL:         url,
			Description: "Endpoint URL",
		},
	},
	}
	c := client.NewAPIClient(&conf)
	tokenStruct, r, err := c.UsersApi.LoginUserWithPwd(ctx).
		LoginRequest(*client.NewLoginRequest(username, password)).Execute()
	if err != nil {
		logger.Warn("could not login to OpenMetadata")
		return nil, err
	}

	r.Body.Close()
	token := fmt.Sprintf("%s %s", tokenStruct.TokenType, tokenStruct.AccessToken)
	conf.DefaultHeader = map[string]string{"Authorization": token}
	return client.NewAPIClient(&conf), nil
}

func (o *OMClient) GetConnectionInformation(ctx context.Context, connectionName string,
	logger hclog.Logger) (*client.DatabaseService, error) {
	c, err := NewOMClient(ctx, logger)
	if err != nil {
		return nil, err
	}

	databaseService, _, err := c.DatabaseServiceApi.GetDatabaseServiceByFQN(ctx, connectionName).Execute()
	if err != nil {
		return nil, err
	}
	return databaseService, nil
}
