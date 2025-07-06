// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso

import (
	"context"
	"log"

	ssh "golang.org/x/crypto/ssh"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/kv-infra/packer-plugin-kubevirt/builder/kubevirt/common"

	"kubevirt.io/client-go/kubecli"
)

type Builder struct {
	config Config
	runner multistep.Runner
	client kubecli.KubevirtClient
}

func (b *Builder) ConfigSpec() hcldec.ObjectSpec {
	return b.config.FlatMapstructure().HCL2Spec()
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, []string, error) {
	warnings, errs := b.config.Prepare(raws...)
	if errs != nil {
		return nil, warnings, errs
	}

	kubeConfig := b.config.KubeConfig
	client, err := kubecli.GetKubevirtClientFromFlags("", kubeConfig)
	if err != nil {
		log.Panicln(err)
	}
	b.client = client
	return nil, warnings, nil
}

func (b *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	steps := []multistep.Step{}
	steps = append(steps,
		&StepCreateIsoVolume{
			config: b.config,
			client: b.client,
		},
		&StepCreateVirtualMachine{
			config: b.config,
			client: b.client,
		},
		&StepBootCommand{
			config: b.config,
			client: b.client,
		},
		&StepWaitInstallation{
			config: b.config,
		},
	)

	if b.config.Communicator == "ssh" {
		sshSteps, err := b.buildSSHSteps()
		if err != nil {
			ui.Errorf("SSH communicator config error: %v", err)
			return nil, nil
		}
		steps = append(steps, sshSteps...)
	}

	steps = append(steps,
		&StepStopVirtualMachine{
			config: b.config,
			client: b.client,
		},
		&StepCreateBootableVolume{
			config: b.config,
			client: b.client,
		},
	)

	state := new(multistep.BasicStateBag)
	state.Put("hook", hook)
	state.Put("ui", ui)

	b.runner = commonsteps.NewRunner(steps, b.config.PackerConfig, ui)
	b.runner.Run(ctx, state)
	return nil, nil
}

func (b *Builder) buildSSHSteps() ([]multistep.Step, []error) {
	commConfig := &communicator.Config{
		Type: b.config.Communicator,
		SSH: communicator.SSH{
			SSHHost:     common.DefaultIPAddress,
			SSHPort:     common.DefaultLocalPort,
			SSHUsername: b.config.SSHUsername,
			SSHPassword: b.config.SSHPassword,
			SSHTimeout:  b.config.SSHWaitTimeout,
		},
	}

	if err := commConfig.Prepare(&interpolate.Context{}); err != nil {
		return nil, err
	}

	steps := []multistep.Step{
		&StepStartPortForward{
			config: b.config,
			client: b.client,
		},
		&communicator.StepConnect{
			Config: commConfig,
			Host: func(state multistep.StateBag) (string, error) {
				return commConfig.SSH.SSHHost, nil
			},
			SSHConfig: func(state multistep.StateBag) (*ssh.ClientConfig, error) {
				return &ssh.ClientConfig{
					User: b.config.SSHUsername,
					Auth: []ssh.AuthMethod{
						ssh.Password(b.config.SSHPassword),
					},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				}, nil
			},
			SSHPort: func(state multistep.StateBag) (int, error) {
				return common.DefaultLocalPort, nil
			},
		},
		&commonsteps.StepProvision{},
	}
	return steps, nil
}
