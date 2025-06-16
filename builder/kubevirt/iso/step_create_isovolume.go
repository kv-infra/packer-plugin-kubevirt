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

type StepCreateIsoVolume struct {
	config Config
	client kubecli.KubevirtClient
}

func (s *StepCreateIsoVolume) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	ui.Say("Creating a new DataVolume that imports ISO from the remote server...")

	name := s.config.Name
	namespace := s.config.Namespace
	isoUrl := s.config.IsoUrl
	isoSize := s.config.IsoSize
	isoVolume := isoVolume(name, isoUrl, isoSize)

	dv, err := s.client.CdiClient().CdiV1beta1().DataVolumes(namespace).Create(ctx, isoVolume, metav1.CreateOptions{})
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if err = waitUntilDataVolumeSucceeded(ctx, s.client, dv); err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepCreateIsoVolume) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	ui.Say("Deleting ISO volume...")

	name := s.config.Name + "-iso"
	namespace := s.config.Namespace

	s.client.CdiClient().CdiV1beta1().DataVolumes(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
