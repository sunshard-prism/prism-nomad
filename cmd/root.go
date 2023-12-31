// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"os"
	"prism/internal/service"

	"github.com/spf13/cobra"
)

var (
	version  = "dev"
	services *service.Service
)

var rootCmd = &cobra.Command{
	Use:     "prism",
	Short:   "Creating a nomad job configuration template.",
	Long:    `Prism creates a nomad job configuration template and deploys it to the cluster.`,
	Version: version,
}

func Execute(service *service.Service) {
	services = service

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("error execute: %s", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
