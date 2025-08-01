#!/bin/bash

week=$(date +%V)
echo "week $week"

if [ $((week%2)) -eq 0 ]; then
	echo "left leg";
else
	echo "right leg";
fi