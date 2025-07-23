// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package iso

import (
	"time"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	KubeConfig              string        `mapstructure:"kube_config"`
	Name                    string        `mapstructure:"name"`
	Namespace               string        `mapstructure:"namespace"`
	IsoVolumeName           string        `mapstructure:"iso_volume_name"`
	DiskSize                string        `mapstructure:"disk_size"`
	InstanceType            string        `mapstructure:"instance_type"`
	InstanceTypeKind        string        `mapstructure:"instance_type_kind"`
	Preference              string        `mapstructure:"preference"`
	PreferenceKind          string        `mapstructure:"preference_kind"`
	MediaFiles              []string      `mapstructure:"media_files"`
	BootCommand             []string      `mapstructure:"boot_command"`
	BootWait                time.Duration `mapstructure:"boot_wait"`
	InstallationWaitTimeout time.Duration `mapstructure:"installation_wait_timeout"`
	Communicator            string        `mapstructure:"communicator"`
	SSHUsername             string        `mapstructure:"ssh_username"`
	SSHPassword             string        `mapstructure:"ssh_password"`
	SSHWaitTimeout          time.Duration `mapstructure:"ssh_wait_timeout"`
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
