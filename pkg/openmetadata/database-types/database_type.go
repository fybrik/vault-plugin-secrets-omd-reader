// Copyright 2023 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package dbtypes

type Databasetype interface {
	ExtractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error)
}
