# Packer Plugin KubeVirt

KubeVirt plugin can be used with HashiCorp Packer to create KubeVirt images.

## Source Build

Run this command from the root directory:

```shell
$ go build -ldflags="-X github.com/codingben/packer-plugin-kubevirt/version.VersionPrerelease=dev" -o packer-plugin-kubevirt
```

To install the compiled plugin, run the following command:
```shell
$ packer plugins install --path packer-plugin-kubevirt github.com/codingben/kubevirt
```
