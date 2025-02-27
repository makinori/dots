# Extras for keyboard

Using CannonKeys Bakeneko60, layout inspired by POK3R

TODO: explain layout i guess

## QMK firmware

-   Clone [QMK](https://docs.qmk.fm/) and setup
-   Save [`keymap.c`](https://raw.githubusercontent.com/makinori/dots/main/keyboard/keymap.c) to `keyboards/cannonkeys/db60/keymaps/maki/keymap.c`
    -   Was [`keymap-socd.c`](https://raw.githubusercontent.com/makinori/dots/main/keyboard/keymap-socd.c) but not using
-   Compile and flash

## Framework 13 with keyd

-   Install `keyd` package
-   Save [`keyd.conf`](https://raw.githubusercontent.com/makinori/dots/main/keyboard/keyd.conf) to `/etc/keyd/default.conf`
-   Enable with `sudo systemd enable --now keyd`
-   Notes:
    -   Top row keys are now **ESC, F1 to F12, Delete**, although **F10** doesn't work
    -   **Mod + Top row delete** will toggle caps lock, just incase
    -   Hardware arrow keys disabled, **left arrow** now mod and **right arrow** now control

---

> <br>
> ⚠️ I'm not using any of the below but kept them in the guide anyway<br>
> <br>

## Linux restore F13 to F24 keys

I don't know, it's not very reliable. But F14 to F19 isn't mapped so anything on X11 I think, so you can always use those.

-   [`fix-fn-keys.sh`](https://raw.githubusercontent.com/makinori/dots/main/keyboard/fix-fn-keys.sh)
-   [`fix-fn-keys-modmap.sh`](https://raw.githubusercontent.com/makinori/dots/main/keyboard/fix-fn-keys-modmap.sh)
-   `yay -S inputplug`<br>
    `inputplug -c /home/maki/git/dots-maki/keyboard/fix-fn-keys.sh`

## Play sound when layer switching

-   Add shortcuts in GNOME for 4 layers with command<br>
    `/home/maki/git/dots-maki/keyboard/play.sh <n>`
