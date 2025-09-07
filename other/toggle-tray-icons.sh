#!/bin/bash

extUUID="appindicatorsupport@rgcjonas.gmail.com"

infoOutput=$(gnome-extensions info $extUUID)

if [[ $infoOutput == *"Enabled: No"* ]]; then
	gnome-extensions enable $extUUID
	echo "Enabled tray icons"
elif [[ $infoOutput == *"Enabled: Yes"* ]]; then
	gnome-extensions disable $extUUID
	echo "Disabled tray icons"
else
	echo "Nothing happened"
fi
