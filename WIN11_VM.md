# Windows VM

There are so many moving parts. This guide is only really for myself.

Really recommend looking through: https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF

Using an AMD Ryzen 7950x

## Files

Current libvirt xml here: https://github.com/makinori/dots/blob/main/win11/win11.xml

Script to start and stop: https://github.com/makinori/dots/blob/main/win11/win11.sh

## Notes

TODO: take everything from xml file apart

-   Definitely disable "Core Isolation Memory Integrity" in Windows Defender

-   `<timer name="hpet" present="yes"/>` for extra performance

## Dynamically Isolate CPUs

https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF#Dynamically_isolating_CPUs

```bash
sudo mkdir -p /etc/libvirt/hooks
sudo vim /etc/libvirt/hooks/qemu
```

```bash
#!/bin/sh

command=$2

if [ "$command" = "started" ]; then
    systemctl set-property --runtime -- system.slice AllowedCPUs=0-7,16-23
    systemctl set-property --runtime -- user.slice AllowedCPUs=0-7,16-23
    systemctl set-property --runtime -- init.scope AllowedCPUs=0-7,16-23
elif [ "$command" = "release" ]; then
    systemctl set-property --runtime -- system.slice AllowedCPUs=0-31
    systemctl set-property --runtime -- user.slice AllowedCPUs=0-31
    systemctl set-property --runtime -- init.scope AllowedCPUs=0-31
fi
```

```bash
sudo chmod +x /etc/libvirt/hooks/qemu
sudo systemctl restart libvirtd.service
```

Check AllowedCPUs using

```bash
systemctl show system.slice | grep AllowedCPUs=

```

## Persistent Evdev

When using `evdev` and `grabToggle="ctrl-ctrl"`, mouse/keyboard input will be lost when they're disconnected. This script proxies them so that they're always available.

Install `yay -S proxydev` from https://gitlab.com/b1gbear/proxydev

`/etc/proxydev/config.toml`

```toml
[[ device ]]
device_type = "keyboard"
vendor_id = 0xca04
product_id = 0xdb62

[[ device ]]
device_type = "mouse"
vendor_id = 0x3554
product_id = 0xf58a
```

```bash
sudo systemctl enable --now proxydev
sudo systemctl status proxydev
```

Now you can use `/dev/input/by-id/proxydev` in libvirt

## Congrats

Now you can chug jug Fortnite gamer on Linux

[chug-jug.webm](https://github.com/user-attachments/assets/a05069fb-6664-42b6-9247-f4d667f52172)

https://www.youtube.com/watch?v=DD_bxaHvt9g
