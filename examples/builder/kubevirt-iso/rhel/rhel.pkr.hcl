# Copyright (c) Red Hat, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    kubevirt = {
      source  = "github.com/kv-infra/kubevirt"
      version = ">= 0.1.0"
    }
  }
}

variable "kube_config" {
  type    = string
  default = "${env("KUBECONFIG")}"
}

source "kubevirt-iso" "rhel" {
  # Kubernetes configuration
  kube_config   = var.kube_config
  name          = "rhel-10-rand-95"
  namespace     = "ben-dev"

  # ISO configuration
  iso_volume_name = "rhel-10-x86-64-iso"

  # VM type and preferences
  disk_size     = "10Gi"
  instance_type = "o1.medium"
  preference    = "rhel.10"

  # Files to include in the ISO installation
  media_files = [
    "./ks.cfg"
  ]

  # Boot process configuration
  # A set of commands to send over VNC connection
  boot_command = [
    "<up>e",                            # Modify GRUB entry
    "<down><down><end>",                # Navigate to kernel line
    " inst.ks=hd:LABEL=OEMDRV:/ks.cfg", # Set kickstart file location
    "<leftCtrlOn>x<leftCtrlOff>"        # Boot with modified command line
  ]
  boot_wait                 = "60s"     # Time to wait after boot starts
  installation_wait_timeout = "10m"     # Timeout for installation to complete

  # SSH configuration
  communicator      = "ssh"
  ssh_username      = "admin"
  ssh_password      = "root"
  ssh_wait_timeout  = "15m"
}

build {
  sources = ["source.kubevirt-iso.rhel"]

  provisioner "shell" {
    inline = [
      "ls -la"
    ]
  }
}
