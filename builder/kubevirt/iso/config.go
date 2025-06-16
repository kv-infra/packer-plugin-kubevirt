// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package iso

import (
	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	KubeConfig   string `mapstructure:"kube_config"`
	IsoUrl       string `mapstructure:"iso_url"`
	IsoSize      string `mapstructure:"iso_size"`
	DiskSize     string `mapstructure:"disk_size"`
	Name         string `mapstructure:"name"`
	Namespace    string `mapstructure:"namespace"`
	InstanceType string `mapstructure:"instance_type"`
	Preference   string `mapstructure:"preference"`
}

func (c *Config) Prepare(raws ...interface{}) ([]string, error) {
	err := config.Decode(c, &config.DecodeOpts{
		PluginType:  "builder.kubevirt.iso",
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, err
	}
	return nil, err
}
