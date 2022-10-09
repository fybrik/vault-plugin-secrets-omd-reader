// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package omsecrets

import (
	"context"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/logical"
)

func getTestBackend(t *testing.T) logical.Backend {
	b, _ := newBackend()

	c := &logical.BackendConfig{
		Logger: hclog.New(&hclog.LoggerOptions{}),
	}
	err := b.Setup(context.Background(), c)
	if err != nil {
		t.Fatalf("unable to create backend: %v", err)
	}
	return b
}

const secretName = "openmetadata-s3" //nolint

func TestGetSecret(t *testing.T) {
	b := getTestBackend(t)

	request := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      secretName,
		Data:      make(map[string]interface{}),
	}

	resp, _ := b.HandleRequest(context.Background(), request)
	if resp.Error() != nil {
		t.Errorf("should not have gotten an error")
	}
}
