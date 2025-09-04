// Copyright (c) Red Hat, Inc.
// SPDX-License-Identifier: MPL-2.0

package iso

type Artifact struct {
	Name string
}

func (a *Artifact) BuilderId() string {
	return "packer.kubevirt.iso"
}

func (a *Artifact) Files() []string {
	return nil
}

func (a *Artifact) Id() string {
	return a.Name
}

func (a *Artifact) String() string {
	return a.Name
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return nil
}
