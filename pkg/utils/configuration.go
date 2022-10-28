// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package utils

import "os"

const (
	OMServerURL string = "OM_SERVER_URL"
	OMUsername  string = "OM_SERVER_USERNAME"
	OMPassword  string = "OM_SERVER_PASSWORD" //nolint:gosec
)

const (
	DefaultOMServerURL string = "http://localhost:8585/api"
	DefaultOMUsername  string = "admin"
	DefaultOMPassword  string = "admin"
)

// GetEnvironmentVariables returns the OM server url, username, password
func GetEnvironmentVariables() (string, string, string) {
	url := os.Getenv(OMServerURL)
	if url == "" {
		url = DefaultOMServerURL
	}
	username := os.Getenv(OMUsername)
	if username == "" {
		username = DefaultOMUsername
	}
	password := os.Getenv(OMPassword)
	if password == "" {
		password = DefaultOMPassword
	}
	return url, username, password
}
