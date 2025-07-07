# Copyright (c) Red Hat, Inc.
# SPDX-License-Identifier: MPL-2.0

name                      = "fedora-42-rand-85"
namespace                 = "ben-dev"
iso_volume_name           = "fedora-42-x86-64-iso"
disk_size                 = "10Gi"
instance_type             = "o1.medium"
preference                = "linux"
media_files               = [
    "./example/kickstarts/ks.cfg"
]
boot_command              = [
    "<up>e",
    "<down><down><end>",
    " inst.ks=hd:LABEL=OEMDRV:/ks.cfg",
    "<leftCtrlOn>x<leftCtrlOff>",
]
boot_wait                 = "60s"
installation_wait_timeout = "10m"
communicator              = "ssh"
ssh_username              = "user"
ssh_password              = "root"
ssh_wait_timeout          = "15m"
