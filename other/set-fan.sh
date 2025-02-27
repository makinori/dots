#!/bin/bash

if [[ $1 = "auto" ]]; then
	value=autofanctrl
	sudo ectool --interface=lpc autofanctrl
else
	sudo ectool --interface=lpc fanduty $1
fi
