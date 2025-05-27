# Maki's Arch Install Guide

> ⚠️ This is never really finished and could get outdated,<br>
> however _"i use arch btw"_ so I'm sure it'll get updated

## 1. Core installation

this guide will only work with uefi. we'll set up partitions first

-   create a 512 MB parition, type: EFI system **(don't if you already have one)**<br>
    create second parition for the rest<br>
    `cfdisk /dev/sda`

-   mkfs FAT32 on first, with label **(only if you made one)**<br>
    `mkfs.fat -F32 /dev/sda1`<br>
    `fatlabel /dev/sda1 MAKI_EFI`

<!--
-   <details>
<summary>ext4 (simple and stable)</summary>
-->

-   choose which root fs you want to use

    would definitely recommend xfs, however it can't be shrinked

    https://fy.blackhats.net.au/blog/2024-08-13-linux-filesystems/

    <details>
    <summary>with encryption</summary>

    -   prepare luks on the other, with label<br>
        `cryptsetup luksFormat /dev/sda2`<br>
        `cryptsetup config /dev/sda2 --label MAKI_ARCH_CRYPT`

    -   mount luks partition<br>
        `cryptsetup open /dev/disk/by-label/MAKI_ARCH_CRYPT root`

    -   use `/dev/mapper/root` instead of `/dev/sda2`

    </details>

    <details>
    <summary>xfs</summary>

    -   mkfs xfs on the other and label<br>
        `mkfs.xfs /dev/sda2`<br>
        `xfs_admin -L MAKI_ARCH /dev/sda2`<br>

    </details>

    <details>
    <summary>ext4</summary>

    -   mkfs ext4 on the other and label<br>
        `mkfs.ext4 /dev/sda2`<br>
        `e2label /dev/sda2 MAKI_ARCH`<br>

    </details>

-   verify labels with<br>
    `lsblk -o name,label`

-   mount the root partition<br>
    `mount /dev/disk/by-label/MAKI_ARCH /mnt`

<!-- </details> -->

<!-- <details>
<summary>btrfs (more complicated, personally unrecommended)</summary>

-   mkfs btrfs on the other<br>
    `mkfs.btrfs /dev/sda2`<br>
    `btrfs filesystem label /dev/sda2 MAKI_ARCH`<br>

-   verify labels with<br>
    `lsblk -o name,label`

-   mount the partition<br>
    `mount /dev/disk/by-label/MAKI_ARCH /mnt`

-   create two subvolumes<br>
    `btrfs sub create /mnt/@`<br>
    `btrfs sub create /mnt/@home`

-   unmount the partition<br>
    `umount /mnt`

-   mount new subvolumes<br>
    `mount -o subvol=@ /dev/disk/by-label/MAKI_ARCH /mnt`<br>
    `mount -o subvol=@home /dev/disk/by-label/MAKI_ARCH /mnt/home`

</details> -->

-   if you **just made** an EFI partition

    -   create directory and mount<br>
        `mkdir /mnt/boot`<br>
        `mount /dev/disk/by-label/MAKI_EFI /mnt/boot`

-   otherwise, if you **already have** one<br>

    -   create seperate directory and mount<br>
        `mkdir /mnt/efi`<br>
        `mount /dev/sda1 /mnt/efi`

-   create a swap file [(see why)](https://chrisdown.name/2018/01/02/in-defence-of-swap.html)<br>
    `mkswap -U clear --size 4G --file /mnt/swapfile`<br>
    `swapon /mnt/swapfile`<br>

-   finally generate fstab<br>
    `mkdir -p /mnt/etc`<br>
    `genfstab -L /mnt > /mnt/etc/fstab`

now we've setup partitions, we'll pacstrap install arch linux and install a boot loader

-   (optional) set geographicly close mirror top<br>
    `nano /etc/pacman.d/mirrorlist`

-   install the base system<br>
    `pacstrap /mnt base base-devel linux linux-firmware linux-headers dhcpcd nano`
    _also install xfsprogs if using xfs_

-   chroot into the new system<br>
    `arch-chroot /mnt`

-   set password for root<br>
    `passwd`

-   install amd-ucode or intel-ucode depending on CPU<br>
    `pacman -S amd-ucode`

-   pick a bootloader, both work with windows. find more here: https://wiki.archlinux.org/title/Arch_boot_process#Boot_loader

    <details>
    <summary>systemd-boot (simple to use, finds windows automatically)</summary>

    -   install systemd bootloader (read --help)<br>
        `bootctl install`

    -   create new boot entry<br>
        `nano /boot/loader/entries/arch.conf`

        ```
        title Arch Linux
        linux /vmlinuz-linux
        initrd /amd-ucode.img
        initrd /initramfs-linux.img
        options root=LABEL=MAKI_ARCH
        options fsck.mode=force fsck.repair=yes
        options rw loglevel=3 nvidia_drm.modeset=1
        ```

        -   if using encryption, update first options line to<br>
            `options cryptdevice=LABEL=MAKI_ARCH_CRYPT:root root=LABEL=MAKI_ARCH`

        -   `nvidia_drm.modeset=1` needed for wayland on nvidia
        -   could remove `loglevel=3` and just set `quiet splash`
        -   could replace `root=LABEL=MAKI_ARCH` with `root=UUID=<uuid>`

        <!--
        -   remove `fsck.mode=force fsck.repair=yes` if using btrfs
        -   add `rootflags=subvol=@` if using btrfs
        -->

    -   set the default entry<br>
        `nano /boot/loader/loader.conf`

        ```
        timeout 5
        default arch.conf
        ```

    </details>

    <details>
    <summary>grub (old and bloated, finds windows with os-prober)</summary>

    <!-- -   _i havent really tried grub with btrfs before. just remove `fsck.mode=force` cause its noop. the rest is up to you_ -->

    -   install a packages<br>
        `pacman -S grub efibootmgr os-prober`

    -   uncomment GRUB_DISABLE_OS_PROBER=false<br>
        `nano /etc/default/grub`

    -   add to end of GRUB_CMDLINE_LINUX_DEFAULT <br>
        `fsck.mode=force nvidia_drm.modeset=1`

    -   if you **just made** an EFI parition<br>
        `mkdir /boot/EFI`<br>
        `grub-install --target=x86_64-efi --efi-directory=/boot/EFI --bootloader-id=GRUB`<br>
        `grub-mkconfig -o /boot/grub/grub.cfg`

    -   if you **already had** an EFI parition<br>
        `grub-install --target=x86_64-efi --efi-directory=/efi --bootloader-id=GRUB`<br>
        `grub-mkconfig -o /boot/grub/grub.cfg` (maybe this needs to be /efi/grub/grub.cfg?)

    </details>

-   if you're encrypting<br>
    `nano /etc/mkinitcpio.conf` add `encrypt` to the end of<br>
    `HOOKS=(... encrypt)` (must be after keymap)<br>
    then run<br>
    `mkinitcpio -P`

-   remove the bootable media and restart<br>
    `exit`<br>
    `reboot`

## 2. Basic settings

-   login as root

-   set hostname<br>
    `echo maki > /etc/hostname`

-   add new user and set password<br>
    `useradd -m -G wheel maki`<br>
    `passwd maki`

-   edit sudoers to allow new user to sudo<br>
    `EDITOR=nano visudo`

    ```
    ## Uncomment to allow members of group wheel to execute any command
    %wheel ALL=(ALL) ALL
    ```

-   check the name of your network interface and start dhcp (don't enable)<br>
    `ip a`<br>
    `systemctl start dhcpcd@enp60s0`<br>

-   uncomment desired locales in locale.gen<br>
    `nano /etc/locale.gen`

-   generate new locales and set one system wide<br>
    `locale-gen`<br>
    `localectl set-locale LANG=en_US.UTF-8`

-   select timezone and set it permanent<br>
    `tzselect`<br>
    `timedatectl set-timezone 'America/Los_Angeles'`

-   set hardware clock and sync using ntp<br>
    `hwclock --systohc --utc`<br>
    `timedatectl set-ntp true`

-   enable systemd-oomd for when we run out-of-memory<br>
    `systemctl enable --now systemd-oomd.service`

-   enable fstrim.timer if you're using an nvme<br>
    `systemctl enable --now fstrim.timer`

-   uncomment color and parallel downloads, and enable x86 repo in pacman.conf<br>
    `nano /etc/pacman.conf`

    ```
    # Misc options
    Color
    ParallelDownloads = 5

    [multilib]
    Include = /etc/pacman.d/mirrorlist
    ```

-   syncronize repositories and reboot<br>
    `pacman -Sy && reboot`

## 3. Desktop environment

-   login as your user

-   check the name of your network interface and start dhcp (don't enable)<br>
    `ip a`<br>
    `systemctl start dhcpcd@enp60s0`<br>

-   install wireplumber for pipewire
    `pacman -S wireplumber pipewire-{pulse,jack,alsa}`<br>
    `systemctl enable --user --now wireplumber pipewire pipewire-pulse`

-   install default video drivers and OpenGL support<br>
    `pacman -S xf86-video-vesa mesa`

-   install video driver specific for your hardware

    see more here: https://wiki.archlinux.org/title/Xorg#Driver_installation

    would recommend for nvidia: https://github.com/Frogging-Family/nvidia-all

    | Brand  | Type        | Packages                                                         |
    | ------ | ----------- | ---------------------------------------------------------------- |
    | NVIDIA | Proprietary | `nvidia-dkms nvidia-utils lib32-nvidia-utils`                    |
    | AMD    | Open Source | `xf86-video-amdgpu lib32-mesa vulkan-radeon lib32-vulkan-radeon` |

-   install xorg server if you dont want to use wayland<br>
    `pacman -S xorg-server`

-   install desktop environment<br>
    (would recommend only installing what you need)<br>
    `pacman -S gnome gnome-extra gnome-themes-extra`

-   may want to stop your dhcp server and install networkmanager for gnome<br>
    `pacman -S networkmanager`<br>
    `systemctl stop dhcpcd@enp60s0`<br>
    `systemctl enable --now NetworkManager`<br>

-   install and enable desktop manager<br>
    `pacman -S gdm`<br>
    `systemctl enable gdm`

-   reboot<br>
    `reboot`

would recommend checking out other guides on [README.md](README.md)
