#!/bin/bash

export MAKEFLAGS="-j32"
export PKGEXT=".pkg.tar" # avoid compression

ignore=()

ignore_args=""

for package in "${ignore[@]}"
do
	ignore_args="$ignore_args --ignore $package"
done

# echo $ignore_args

sudo pacman -Syu $ignore_args

if [[ "$1" != "--pacman" ]]; then
	# yay -Syu --noconfirm $ignore_args
	yay -Syu $ignore_args
fi
