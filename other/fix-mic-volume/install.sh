#!/bin/bash

SCRIPT_DIR=$(realpath "$(dirname "$0")")
cd $SCRIPT_DIR

CGO_ENABLED=0 GOOS=linux go build -o fix-mic-volume

SERVICE_PATH=~/.config/systemd/user/fix-mic-volume.service
INSTALL_PATH=$SCRIPT_DIR/fix-mic-volume

rm -f $SERVICE_PATH
cp fix-mic-volume.service $SERVICE_PATH

sed -i "s#{{.installPath}}#"$INSTALL_PATH"#g" $SERVICE_PATH

systemctl --user daemon-reload
systemctl --user enable --now fix-mic-volume.service

systemctl --user status fix-mic-volume.service