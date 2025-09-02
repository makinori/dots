#!/bin/bash

set -e

SCRIPT_DIR=$(realpath "$(dirname "$0")")
cd $SCRIPT_DIR

export CGO_ENABLED=0
export GOOS=linux

if [[ $@ == *"--laptop"* ]]; then
	echo "building for laptop"
	go build -o maki-audio-helper -tags laptop
else
	echo "building for desktop"
	go build -o maki-audio-helper
fi

SERVICE_PATH=~/.config/systemd/user/maki-audio-helper.service
INSTALL_PATH=$SCRIPT_DIR/maki-audio-helper

rm -f $SERVICE_PATH
cp maki-audio-helper.service $SERVICE_PATH

sed -i "s#{{.installPath}}#"$INSTALL_PATH"#g" $SERVICE_PATH

systemctl --user daemon-reload
systemctl --user enable maki-audio-helper.service
systemctl --user restart --now maki-audio-helper.service

sleep 0.5

systemctl --user status maki-audio-helper.service