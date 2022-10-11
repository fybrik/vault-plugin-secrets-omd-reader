// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package utils

import "os"

const (
	OMServerURL        string = "OM_SERVER_URL"
	DefaultOMServerURL string = "http://localhost:8585/api"
)

// GetOMServerURL returns the OM server url
func GetOMServerURL() string {
	url := os.Getenv(OMServerURL)
	if url == "" {
		url = DefaultOMServerURL
	}
	return url
}
