<!--
  Include a short overview about the plugin.

  This document is a great location for creating a table of contents for each
  of the components the plugin may provide. This document should load automatically
  when navigating to the docs directory for a plugin.

-->

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    name = {
      source  = "github.com/hashicorp/kubevirt"
      version = ">=0.0.1"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
$ packer plugins install github.com/hashicorp/kubevirt
```

### Components

#### Builders

- [kubevirt-clone](/packer/integrations/hashicorp/kubevirt/latest/components/builder/kubevirt-clone) - This builder clones a virtual machine from an existing template and then modifies and saves it as a new template.
