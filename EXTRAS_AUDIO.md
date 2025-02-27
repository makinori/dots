# Extras for audio

Everything here has been tested with PipeWire

## Force microphone to fixed volume

Uses wireplumber and systemd

-   ```bash
    wget https://raw.githubusercontent.com/makinori/dots/main/other/fix-mic-volume.lua
    chmod +x fix-mic-volume.lua
    ```

-   update variables at the top of the lua file

-   `vim ~/.config/systemd/user/fix-mic-volume.service`

    ```service
    [Unit]
    Description=Fix Mic Volume
    After=pipewire.service

    [Service]
    # ExecStartPre=/bin/sleep 2
    ExecStart=/home/maki/fix-mic-volume.lua
    Type=simple
    # treat as a lowest priority program
    Nice=19

    [Install]
    WantedBy=default.target
    ```

-   `systemctl enable --now --user fix-mic-volume.service`
-   `systemctl status --user fix-mic-volume.service` make sure its not erroring

## Virtual audio cable

After installing the service file, create virtual audio cables using:

`systemctl enable --user --now virtual-cable@OBS`<br>
(OBS as example for monitoring)

`vim ~/.config/systemd/user/virtual-cable@.service`

```service
[Unit]
Description=%i Virtual Cable
After=pipewire-pulse.service

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/bin/pactl load-module module-null-sink \
sink_name=%i \
sink_properties="'device.description=\"%i Virtual Cable\"'"
ExecStart=/usr/bin/pactl load-module module-remap-source \
source_name=%i master=%i.monitor \
source_properties="'device.description=\"%i Virtual Cable\"'"
ExecStop=bash -c "/usr/bin/pactl unload-module \
$(pactl list short modules | grep sink_name=%i | cut -d$'\t' -f1)"
ExecStop=bash -c "/usr/bin/pactl unload-module \
$(pactl list short modules | grep source_name=%i | cut -d$'\t' -f1)"

[Install]
WantedBy=default.target
```

# Setup DeaDBeeF

-   `yay -S deadbeef-git deadbeef-mpris2-plugin deadbeef-plugin-fb-gtk3-git deadbeef-plugin-spectrogram-gtk3-git deadbeef-plugin-discord-git`
-   [`.config/deadbeef/config`](https://raw.githubusercontent.com/makinori/dots/main/.config/deadbeef/config)
