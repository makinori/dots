#!/bin/bash

cd "$(dirname "$0")"

export STEAM_COMPAT_DATA_PATH=$(pwd)/compat
export STEAM_COMPAT_CLIENT_INSTALL_PATH=$HOME/.steam/steam

export PROTON_DIR=/usr/share/steam/compatibilitytools.d/proton-ge-custom
export PROTON=$PROTON_DIR/proton

export WINEPREFIX=$STEAM_COMPAT_DATA_PATH/pfx 
export WINE=$PROTON_DIR/files/bin/wine64
export WINEARCH=win64

mkdir -p $STEAM_COMPAT_DATA_PATH
touch $STEAM_COMPAT_DATA_PATH/pfx.lock
touch $STEAM_COMPAT_DATA_PATH/tracked_files

# $WINE $@
# winetricks

$PROTON run "path to exe"
