# Packer Plugin KubeVirt

The `KubeVirt` plugin can be used with [HashiCorp Packer](https://www.packer.io) to create KubeVirt images.

## Components

### Builders

- `kubevirt-iso` - This builder starts from a ISO file and builds virtual machine image on a KubeVirt cluster.

## Installation

### Source Build

Clone this repository and run build from the root directory:

```shell
$ go build -ldflags="-X github.com/kubevirt-infra/packer-plugin-kubevirt/version.Version=0.0.1" -o packer-plugin-kubevirt
```

To install the compiled plugin, run the following command:

```shell
$ packer plugins install --path packer-plugin-kubevirt github.com/kubevirt-infra/kubevirt
```
