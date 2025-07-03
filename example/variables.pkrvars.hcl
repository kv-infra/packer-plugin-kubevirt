# Copyright (c) Red Hat, Inc.
# SPDX-License-Identifier: MPL-2.0

name             = "fedora42"
namespace        = "ben-dev"
iso_volume_name  = "fedora42-x86-64-iso"
disk_size        = "10Gi"
instance_type    = "o1.medium"
preference       = "linux"
boot_command     = [
    "<up>e",
    "<down><down><end>",
    " inst.ks=hd:LABEL=OEMDRV:/ks.cfg",
    "<leftCtrlOn>x<leftCtrlOff>",
]
boot_wait        = "60s"
communicator     = "ssh"
ssh_username     = "user"
ssh_password     = "root"
ssh_wait_timeout = "15m"
