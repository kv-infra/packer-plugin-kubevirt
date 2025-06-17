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

variable "iso_url" {
  type = string
}

variable "iso_size" {
  type = string
}

variable "disk_size" {
  type = string
}

variable "name" {
  type = string
}

variable "namespace" {
  type = string
}

variable "instance_type" {
  type = string
}

variable "preference" {
  type = string
}

source "kubevirt-iso" "example" {
  kube_config     = var.kube_config
  iso_url         = var.iso_url
  iso_size        = var.iso_size
  disk_size       = var.disk_size
  name            = var.name
  namespace       = var.namespace
  instance_type   = var.instance_type
  preference      = var.preference
}

build {
  sources = ["source.kubevirt-iso.example"]
}
