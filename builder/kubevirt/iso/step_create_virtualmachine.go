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
	ui.Say("Creating a new temporary VirutalMachine...")

	name := s.config.Name
	namespace := s.config.Namespace
	isoVolumeName := s.config.IsoVolumeName
	diskSize := s.config.DiskSize
	instanceTypeName := s.config.InstanceType
	preferenceName := s.config.Preference
	v1VirtualMachine := virtualMachine(name, isoVolumeName, diskSize, instanceTypeName, preferenceName)

	_, err := s.client.VirtualMachine(namespace).Create(ctx, v1VirtualMachine, metav1.CreateOptions{})
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepCreateVirtualMachine) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	ui.Say("Deleting VirutalMachine...")

	name := s.config.Name + "-vm"
	namespace := s.config.Namespace

	s.client.VirtualMachine(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
