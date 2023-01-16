// Copyright 2023 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package dbtypes

import "encoding/json"

const FybrikAccessKeyString = "access_key"
const FybrikSecretKeyString = "secret_key"

type S3Config struct {
	ConfigSource struct {
		SecurityConfig struct {
			AccessKey string `json:"awsAccessKeyId"`
			SecretKey string `json:"awsSecretAccessKey"`
		} `json:"securityConfig"`
	} `json:"configSource"`
}

type S3 struct {
}

func (s S3) ExtractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error) {
	var c S3Config
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
