// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/kubevirt-infra/packer-plugin-kubevirt/version"
)

func main() {
	setup := plugin.NewSet()
	setup.SetVersion(version.PluginVersion)

	if err := setup.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
