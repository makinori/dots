# Extras for audio

Everything here has been tested with PipeWire

## Force microphone to fixed volume

Uses Pulse library in Go, please use with PipeWire

-   git clone this repo

-   check and modify `./programs/fix-mic-volume`

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

## Low latency high quality Schiit Modi 3+

-   `mkdir -p ~/.config/wireplumber/wireplumber.conf.d`
-   Save [`51-schiit-modi.conf`](https://github.com/makinori/dots/tree/main/.config/wireplumber/wireplumber.conf.d/51-schiit-modi.conf)

## Setup DeaDBeeF

-   `yay -S deadbeef-git deadbeef-mpris2-plugin deadbeef-plugin-fb-gtk3-git deadbeef-plugin-spectrogram-gtk3-git`
-   [`.config/deadbeef`](https://github.com/makinori/dots/tree/main/.config/deadbeef)

<!-- deadbeef-plugin-waveform-gtk3-git -->

## Setup fooyin

TODO: make spectrum and spectogram plugin and add album art support to directory browser

-   `yay -S fooyin`
-   [`.config/fooyin`](https://github.com/makinori/dots/tree/main/.config/fooyin)
