// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"

	omsecrets "fybrik.io/vault-plugin-secrets-omd-reader/pkg/omsecrets"
)

const ErrorParseArguments = 1
const ErrorServerPlugin = 2

func main() {
	// Boilerplate code to get started.
	// Please see https://www.hashicorp.com/blog/building-a-vault-secure-plugin for more info.
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	err := flags.Parse(os.Args[1:]) //nolint //temporarily needed due to golangci-lint bug: https://github.com/golangci/golangci-lint/discussions/3286
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("Error in flags.Parse", err)
		os.Exit(ErrorParseArguments)
	}

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: omsecrets.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(ErrorServerPlugin)
	}
}
