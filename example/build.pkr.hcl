# Copyright (c) Red Hat, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    kubevirt = {
      source  = "github.com/kv-infra/kubevirt"
      version = ">= 0.0.1"
    }
  }
}

variable "kube_config" {
  type = string
  default = "${env("KUBECONFIG")}"
}

variable "name" {
  type = string
}

variable "namespace" {
  type = string
}

variable "iso_volume_name" {
  type = string
}

variable "disk_size" {
  type = string
}

variable "instance_type" {
  type = string
}

variable "preference" {
  type = string
}

variable "media_files" {
  type = list(string)
}

variable "boot_command" {
  type = list(string)
}

variable "boot_wait" {
  type = string
}

variable "installation_wait_timeout" {
  type = string
}

variable "communicator" {
  type = string
}

variable "ssh_username" {
  type = string
}

variable "ssh_password" {
  type = string
}

variable "ssh_wait_timeout" {
  type = string
}

source "kubevirt-iso" "example" {
  kube_config                  = var.kube_config
  name                         = var.name
  namespace                    = var.namespace
  iso_volume_name              = var.iso_volume_name
  disk_size                    = var.disk_size
  instance_type                = var.instance_type
  preference                   = var.preference
  media_files                  = var.media_files
  boot_command                 = var.boot_command
  boot_wait                    = var.boot_wait
  installation_wait_timeout    = var.installation_wait_timeout
  communicator                 = var.communicator
  ssh_username                 = var.ssh_username
  ssh_password                 = var.ssh_password
  ssh_wait_timeout             = var.ssh_wait_timeout
}

build {
  sources = ["source.kubevirt-iso.example"]

  provisioner "shell" {
    inline = [
      "ls -la"
    ]
  }
}
