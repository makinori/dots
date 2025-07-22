#!/bin/bash

if [[ $1 == 0 ]]; then
	sudo find /etc/pam.d -type f -exec \
	sed -i -E "s/^([^#].+\spam_fprintd.so)/#\1/g" "{}" \;
	echo "disabled"

elif [[ $1 == 1 ]]; then
	sudo find /etc/pam.d -type f -exec \
	sed -i -E "s/^#(.+\spam_fprintd.so)/\1/g" "{}" \;
	echo "enabled"

else
	echo "usage: 0 to disable, 1 to enable"

fi