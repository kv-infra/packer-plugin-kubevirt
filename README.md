# Packer Plugin KubeVirt

The `KubeVirt` plugin can be used with [HashiCorp Packer](https://www.packer.io) to create KubeVirt images.

**Note**: This plugin is under development and is not production ready.

## Packer

[Packer](https://developer.hashicorp.com/packer) is a tool for creating identical machine images from a single source template.

To get started, see the [Packer installation guide](https://developer.hashicorp.com/packer/install).

## Plugin Features

- **HCL Templating** – Use HashiCorp Configuration Language (HCL2) for defining infrastructure as code.
- **ISO-based VM Creation** – Build VM images from ISO using the `kubevirt-iso` builder.
- **ISO Media Files** – Embed additional files into installation (e.g., `kickstart.cfg`).
- **Automated Boot Configuration** – Automate the VM boot process using a set of commands over VNC.
- **Integrated SSH Access** – Enable VM provisioning and customization over SSH.

## Components

### Builders

- `kubevirt-iso` - This builder starts from a ISO file and builds virtual machine image on a KubeVirt cluster.

## Prerequisites

- [Packer](https://packer.io)
- [Kubernetes](https://kubernetes.io) with [KubeVirt](https://kubevirt.io) installed

## Installation

### Using Released Binary

Download the latest release from the [Releases](https://github.com/kv-infra/packer-plugin-kubevirt/releases) page and then install the plugin:

```shell
$ packer plugins install --path packer-plugin-kubevirt github.com/kv-infra/kubevirt
```

### Building From Source

Clone the repository and build the plugin from the root directory:

```shell
$ go build -ldflags="-X github.com/kv-infra/packer-plugin-kubevirt/version.Version=0.1.0" -o packer-plugin-kubevirt
```

Then install the compiled plugin:

```shell
$ packer plugins install --path packer-plugin-kubevirt github.com/kv-infra/kubevirt
```

## Usage

Refer to the usage guidance in the [examples](./examples/builder/kubevirt-iso) of this plugin.
