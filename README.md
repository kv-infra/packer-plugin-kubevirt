# Packer Plugin KubeVirt

The `KubeVirt` plugin can be used with [HashiCorp Packer](https://www.packer.io) to create KubeVirt images.

## Components

### Builders

- `kubevirt-iso` - This builder starts from a ISO file and builds virtual machine image on a KubeVirt cluster.

## Requirements

- [KubeVirt](https://kubevirt.io)
- [Kubernetes](https://kubernetes.io)

## Installation

### Source Build

Clone this repository and run build from the root directory:

```shell
$ go build -ldflags="-X github.com/kv-infra/packer-plugin-kubevirt/version.Version=0.0.1" -o packer-plugin-kubevirt
```

To install the compiled plugin, run the following command:

```shell
$ packer plugins install --path packer-plugin-kubevirt github.com/kv-infra/kubevirt
```

## Usage

> 💡 Ensure you are logged in to the Kubernetes cluster and that KubeVirt is installed.

Export variable below that is used by the Packer builder:

```shell
$ export KUBECONFIG=~/.kube/config
```

Deploy a DataVolume to import the Fedora 42 ISO:

```shell
$ kubectl apply -f ./example/fedora42-x86-64-iso.yaml
```

Create a ConfigMap containing the kickstart file for automated installation:

```shell
$ kubectl create configmap oemdrv-cm --from-file=./example/ks.cfg
```

Run the Packer build with custom variables:

```shell
$ packer build -var-file=./example/variables.pkrvars.hcl ./example/build.pkr.hcl
```
