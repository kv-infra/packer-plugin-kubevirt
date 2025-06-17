// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso

import (
	"context"
	"log"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"

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
