#!/bin/bash

# please actually support artists on bandcamp
# im unfortunately not able to sometimes
# but i absolutely support artists when i can

set -euo pipefail

inputURL=$1
extraArgs=${*:2}

if [[ -z $inputURL ]]; then
	echo "usage: <url>"
	exit 1
fi

# python -v venv ~/venv
source ~/venv/bin/activate
# pip install -U bandcamp-downloader

mkdir -p ~/bandcamp-dl

args=(
	# --embed-art # just makes everything bigger. only do for tracks not albums
	# --no-slugify
	--keep-spaces
	--keep-upper
	--base-dir ~/bandcamp-dl
	--no-confirm
	$extraArgs
)

args=${args[@]}


if [[ $inputURL == *"/album/"* ]]; then
	bandcamp-dl $args \
	--template "%{artist} - %{album}/%{track} - %{title}" \
	$inputURL
elif [[ $inputURL == *"/track/"* ]]; then
	bandcamp-dl $args \
	--embed-art --full-album \
	--template "%{artist} - %{album}" \
	$inputURL
elif [[ $inputURL == *".bandcamp.com" || $inputURL == *".bandcamp.com/" ]]; then
	# trim /
	baseURL=${inputURL%/}

	# TODO: only downloads first page
	curl -o ~/bandcamp-dl.sh.tmp $baseURL

	while IFS= read -r url; do
		if [[ $url == "/"* ]]; then
			url=$baseURL$url
		fi

		echo "downloading: $url"

		# invoke self
		$0 $url

	done < <(
		xmllint --html --xpath "//a/@href" ~/bandcamp-dl.sh.tmp 2>/dev/null | \
		tr -s " " "\n" | \
		sed -E 's/href="([^"]*)"/\1/g' | \
		grep -E '/(track|album)/'
	)

	rm -f ~/bandcamp-dl.sh.tmp

	echo "please check everything was downloaded"
else
	echo "unsure what url this is"
fi
