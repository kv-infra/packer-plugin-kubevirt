# Packer Plugin KubeVirt

The `KubeVirt` plugin can be used with [HashiCorp Packer](https://www.packer.io) to create KubeVirt images.

## Components

### Builders

- `kubevirt-iso` - This builder starts from a ISO file and builds virtual machine image on a KubeVirt cluster.

## Requirements

- [KubeVirt](https://kubevirt.io)
- [Kubernetes](https://kubernetes.io)

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

> 💡 Ensure you are logged in to the Kubernetes cluster and that KubeVirt is installed.

1. Export variable below that is used by the Packer builder:

```shell
$ export KUBECONFIG=~/.kube/config
```

2. Deploy a DataVolume to import the Fedora 42 ISO:

```shell
$ kubectl apply -f ./example/iso/fedora42-x86-64-iso.yaml
```

3. Run the Packer build with custom variables:

```shell
$ packer build -var-file=./example/variables.pkrvars.hcl ./example/build.pkr.hcl
```
