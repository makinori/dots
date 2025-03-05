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

Original: https://github.com/aiberia/persistent-evdev

Find modfied files in `win11` folder

```bash
sudo pacman -Sy python-evdev python-pyudev
sudo ./install.sh
sudo systemctl restart systemd-udevd # apply rules so we dont get double input
# test first, then close
sudo /opt/persistent-evdev/persistent-evdev.py /opt/persistent-evdev/config.json
# enable and start service
sudo systemctl enable --now persistent-evdev.service
sudo systemctl status persistent-evdev.service
```

Now you can use the following in libvirt

```
/dev/input/by-id/uinput-persist-keyboard
/dev/input/by-id/uinput-persist-mouse
```

> **⚠ Note:** there's a bug where the script can leak gigabytes of memory

## Congrats

Now you can chug jug Fortnite gamer on Linux

[chug-jug.webm](https://github.com/user-attachments/assets/a05069fb-6664-42b6-9247-f4d667f52172)

https://www.youtube.com/watch?v=DD_bxaHvt9g
