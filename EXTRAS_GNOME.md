# Extras for GNOME

## Recommended settings

-   use fixed number of workspaces to 8 in **Settings > Multitasking > Workspaces**

-   update keybinds in **Settings > Keyboard > View and Customize Shortcuts**

    > using <kbd>alt</kbd> here cause ended up getting bad rsi with <kbd>super</kbd>

    -   under `Naviation`

        -   `Move window one workspace to the left` to **Shift + Alt + Q**
        -   `Move window one workspace to the right` to **Shift + Alt + E**
        -   `Switch applications` to **Disabled**
        -   `Switch to workspace on the left` to **Alt + Q**
        -   `Switch to workspace on the right` to **Alt + E**
        -   ❗ `Switch windows` to **Alt + Tab**

    -   under `System`

        -   `Show all apps` to **Disabled**
        -   `Show the run command prompty` to **Alt + F12**

    -   under `Custom`
        -   `Open terminal` with `gnome-terminal` to **Alt + Return**
        -   `Playerctl next` with `playerctl next` to **Audio next**
        -   `Playerctl play/pause` with `playerctl play-pause` to **Audio play**
        -   `Playerctl previous` with `playerctl previous` to **Audio previous**

-   disable alt menu

    -   gtk

        ```ini
        [Settings]
        gtk-enable-mnemonics=0
        gtk-auto-mnemonics=0
        ```

        `~/.config/gtk-2.0/settings.ini`</br>
        `~/.config/gtk-3.0/settings.ini`</br>
        `~/.config/gtk-4.0/settings.ini`</br>

    -   vscode

        ```json
        "window.titleBarStyle": "custom",
        "window.customMenuBarAltFocus": false
        ```

-   using tweaks, disable **Middle Click Paste** in **Keyboard & Mouse**

-   if using 800 dpi for mouse, half sensitivity and disable mouse acceleration<br>
    `gsettings set org.gnome.desktop.peripherals.mouse speed -0.5`<br>
    `gsettings set org.gnome.desktop.peripherals.mouse accel-profile flat`

-   resize windows using mod+rmb<br>
    `gsettings set org.gnome.desktop.wm.preferences resize-with-right-button true`

-   incase `file://` uris dont open in gnome nautilus<br>
    `xdg-mime default org.gnome.Nautilus.desktop inode/directory`

-   remove arch logo from gdm<br>
    `sudo -u gdm dbus-launch gsettings set org.gnome.login-screen logo ""`

## Themes

-   please dont theme gnome https://stopthemingmy.app
-   but do use [Papirus Icon Theme](https://github.com/PapirusDevelopmentTeam/papirus-icon-theme)
-   change font to [SN Pro](https://supernotes.app/open-source/sn-pro/) `yay -S ots-sn-pro`

## Nautilus

If using VSCode or VSCodium, add button to open folders

`yay -S --noconfirm code-nautilus-git`

Would recommend `nautilus-code-git` but I like to<br/>
`ln -s /usr/bin/codium /usr/bin/code` which above will show only once

## Extensions

recommended to install below using yay cause devs arent updating their extensions on gnome.org

`yay -S --noconfirm gnome-shell-extension-appindicator gnome-shell-extension-arch-update gnome-shell-extension-hidetopbar-git`

-   [AppIndicator and KStatusNotifierItem Support](https://extensions.gnome.org/extension/615/appindicator-support/) adds tray icons back
-   [Arch Update Indicator](https://extensions.gnome.org/extension/1010/archlinux-updates-indicator/) shows package updates
-   [Hide Top Bar](https://extensions.gnome.org/extension/545/hide-top-bar/) gives back space

    -   In advanced settings, change command to check for package updates:<br>
        `/bin/sh -c "/usr/bin/checkupdates && /usr/bin/yay -Qua"`
    -   Command to update packages:<br>
        `gnome-terminal -- /bin/sh -c "~/update-all.sh; echo Done - Press enter to exit; read _"`
    -   In home folder, place [update-all.sh](https://raw.githubusercontent.com/makinori/dots/main/other/update-all.sh)

extras but not using. the less addons the better

-   [Freon](https://extensions.gnome.org/extension/841/freon/) shows cpu temp and more
-   [GameMode](https://extensions.gnome.org/extension/1852/gamemode/) shows indicator when `gamemoderun`
-   [Rounded Window Corners](https://extensions.gnome.org/extension/5237/rounded-window-corners/) is just nice

honorable but not using

-   [Dash to Panel](https://extensions.gnome.org/extension/1160/dash-to-panel/)
-   [Dash to Dock](https://extensions.gnome.org/extension/307/dash-to-dock/)
-   [Vertical overview](https://extensions.gnome.org/extension/4144/vertical-overview/)
-   [Grand Theft Focus](https://extensions.gnome.org/extension/5410/grand-theft-focus/) brings ready windows into focus
