// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso

import (
	"context"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kubevirt.io/client-go/kubecli"
)

type StepCreateVirtualMachine struct {
	config Config
	client kubecli.KubevirtClient
}

func (s *StepCreateVirtualMachine) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	name := s.config.Name
	namespace := s.config.Namespace
	isoVolumeName := s.config.IsoVolumeName
	diskSize := s.config.DiskSize
	instanceTypeName := s.config.InstanceType
	preferenceName := s.config.Preference
	virtualMachine := virtualMachine(name, isoVolumeName, diskSize, instanceTypeName, preferenceName)

	ui.Sayf("Creating a new temporary VirutalMachine (%s/%s)...", namespace, name)

	_, err := s.client.VirtualMachine(namespace).Create(ctx, virtualMachine, metav1.CreateOptions{})
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepCreateVirtualMachine) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	name := s.config.Name
	namespace := s.config.Namespace

	ui.Sayf("Deleting VirutalMachine (%s/%s)...", namespace, name)

	s.client.VirtualMachine(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
