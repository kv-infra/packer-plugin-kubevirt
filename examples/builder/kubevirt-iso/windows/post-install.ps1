# Install QEMU Tools (Drivers)
Start-Process msiexec -Wait -ArgumentList "/i E:\virtio-win-gt-x64.msi /qn /passive /norestart"

# Install QEMU Guest Agent
Start-Process msiexec -Wait -ArgumentList "/i E:\guest-agent\qemu-ga-x86_64.msi /qn /passive /norestart"

# Rename cached unattend.xml to avoid it is picked up by sysprep
mv C:\Windows\Panther\unattend.xml C:\Windows\Panther\unattend.install.xml

# Eject CD, to avoid that the autounattend.xml on the CD is picked up by sysprep
(New-Object -COMObject Shell.Application).NameSpace(17).ParseName("F:").InvokeVerb("Eject")

# Run Sysprep and Shutdown
C:\Windows\System32\Sysprep\sysprep.exe /generalize /oobe /shutdown /mode:vm
