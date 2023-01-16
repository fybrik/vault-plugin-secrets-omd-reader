// Copyright 2023 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package dbtypes

import "encoding/json"

const Username = "username"
const Password = "password"

type MysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mysql struct {
}

func (m Mysql) ExtractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error) {
	var c MysqlConfig
	configBytes, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configBytes, &c)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		Username: c.Username,
		Password: c.Password,
	}, nil
}
