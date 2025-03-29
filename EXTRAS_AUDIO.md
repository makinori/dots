# Extras for audio

Everything here has been tested with PipeWire

## Force microphone to fixed volume

Uses Pulse library in Go, please use with PipeWire

-   git clone this repo

-   check and modify `./other/fix-mic-volume`

-   run `./install.sh`

-   `systemctl status --user fix-mic-volume.service`

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
