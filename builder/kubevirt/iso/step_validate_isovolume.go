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
	isoVolumeNamespace := s.config.Namespace
	isoVolumeName := s.config.IsoVolumeName

	ui.Sayf("Validating existence of the ISO volume %s in %s namespace...", isoVolumeName, isoVolumeNamespace)

	_, err := s.client.CdiClient().CdiV1beta1().DataVolumes(isoVolumeNamespace).Get(ctx, isoVolumeName, metav1.GetOptions{})
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Sayf("ISO volume %s is expected to be in Succeeded phase...", isoVolumeName)

	if err := waitUntilDataVolumeSucceeded(ctx, s.client, isoVolumeNamespace, isoVolumeName); err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *StepCreateIsoVolume) Cleanup(state multistep.StateBag) {
	// Left blank intentionally
}
