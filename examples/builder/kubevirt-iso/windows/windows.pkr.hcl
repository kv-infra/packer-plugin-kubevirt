# Copyright (c) Red Hat, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    kubevirt = {
      source  = "github.com/kv-infra/kubevirt"
      version = ">= 0.4.0"
    }
  }
}

variable "kube_config" {
  type    = string
  default = "${env("KUBECONFIG")}"
}

source "kubevirt-iso" "windows" {
  # Kubernetes configuration
  kube_config   = var.kube_config
  name          = "windows-11-rand-185"
  namespace     = "ben-dev"

  # ISO configuration
  iso_volume_name = "windows-11-x86-64-iso"

  # VM type and preferences
  disk_size     = "64Gi"
  instance_type = "u1.large"
  preference    = "windows.11.virtio"

  # Files to include in the ISO installation
  media_files = [
    "./autounattend.xml",
    "./post-install.ps1"
  ]

  # Boot process configuration
  # A set of commands to send over VNC connection
  boot_command = [
    "<spacebar><wait>",                # Bypass press any key press challenge
  ]
  boot_wait                 = "5s"     # Time to wait after boot starts
  installation_wait_timeout = "20m"    # Timeout for installation to complete
}

build {
  sources = ["source.kubevirt-iso.windows"]
}
