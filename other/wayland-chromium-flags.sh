#!/bin/bash

function write_config {
	file_path="$HOME/.config/$1-flags.conf"

	echo "$file_path"

	cat << EOF > "$file_path"
--ozone-platform=wayland
--enable-features=UseOzonePlatform,WaylandWindowDecorations
--enable-webrtc-pipewire-capturer
EOF
}

# TODO: do we need so many electrons?

write_config chrome
write_config chromium
write_config code
write_config codium
write_config electron25
write_config electron26
write_config electron27
write_config electron28
write_config electron29
write_config electron30
write_config electron31
write_config electron32
write_config electron33
write_config electron34
write_config electron
