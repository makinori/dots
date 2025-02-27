# Extras full of goodies

-   install yay aur helper

    ```
    git clone https://aur.archlinux.org/yay.git
    cd yay
    makepkg -si
    ```

-   firefox with vertical tabs and more

    <!-- -   install mercury browser<br />
        `yay -S mercury-browser-avx2-bin` -->

    -   install [sideberry](https://addons.mozilla.org/en-US/firefox/addon/sidebery/) extension
    -   css hack [to hide native tabs](<https://github.com/mbnuqw/sidebery/wiki/Firefox-Styles-Snippets-(via-userChrome.css)#completely-hide-native-tabs-strip>)
    -   also add `#sidebar-header { display: none; }`
    -   some custom favicon too, entire file here: [`firefox-userChrome.css`](https://github.com/makinori/dots/blob/main/other/firefox-userChrome.css)

    -   install [`minimalist-dracula-darker.xpi`](https://github.com/makinori/dots/blob/main/other/minimalist-dracula-darker.xpi) theme<br />
        (modified from [MinimalistFox](https://github.com/canbeardig/MinimalistFox))

-   make emojis work by running<br />
    `yay -S noto-fonts-emoji noto-color-emoji-fontconfig`

-   disable `*-debug` packages

    -   edit `/etc/makepkg.conf`
    -   modify `OPTIONS=(... debug)` to `OPTIONS=(... !debug)`

-   fix file permissions when copying from ntfs

    ```bash
    find Files/* -type d -exec chmod 755 {} +
    find Files/* -type f -exec chmod 644 {} +
    ```

-   fix chromium and electron on wayland

    ```bash
    ./other/wayland-chromium-flags.sh
    ```

## Not really using these

-   when using grub2win (try to avoid it i guess)

    ```
    insmod all_video
    set root='(hd0,4)'
    linux /boot/vmlinuz-linux root=/dev/nvme0n1p4 rw
    initrd /boot/initramfs-linux.img
    ```

-   setup alacritty

    [`.config/alacritty/alacritty.yml`](https://raw.githubusercontent.com/makinori/dots/main/.config/alacritty/alacritty.yml) (haven't really touched this in a while)

-   setup smb for sonos music library

    -   `sudo pacman -S samba`
    -   [`/etc/samba/smb.conf`](https://raw.githubusercontent.com/makinori/dots/main/etc/samba/smb.conf)
    -   `sudo systemctl enable smb && sudo systemctl start smb`

-   guide to install affinity suite

    https://codeberg.org/wanesty/affinity-wine-docs/src/branch/guide-wine8.14

-   use [quickemu](https://aur.archlinux.org/packages/quickemu) for easy windows or mac vm

    > im currently doing a gpu passthrough which is quite a bit more complicated.<br>
    > find more info here https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF<br>
    > i also highly recommend https://looking-glass.io

    -   `mkdir ~/quickemu && cd ~/quickemu`
    -   `quickget windows 11`
    -   `quickemu --vm windows-11.conf`
    -   install spice tools from attached drive and shutdown vm
    -   download and run [`run-windows.sh`](https://github.com/makinori/dots/blob/main/other/run-windows.sh)
    -   install software gl+dx support https://github.com/pal1000/mesa-dist-win/releases
        -   in admin run `systemwidedeploy.cmd` and `1. Core desktop OpenGL drivers`
        -   test with [GPU Caps Viewer](https://www.geeks3d.com/dlz/). should be higher than gl 1.0
