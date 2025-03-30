#!/bin/bash

set -e

SCRIPT_DIR=$(realpath "$(dirname "$0")")
cd $SCRIPT_DIR

CGO_ENABLED=0 GOOS=linux go build -o ~/maki-immich

echo "installed to ~/maki-immich"
echo "sample config for nautilus:"
echo ""
echo "~/.local/share/actions-for-nautilus/config.json"
echo ""
echo "{
	"debug": false,
	"actions": [
		{
			"type": "command",
			"label": "Maki Immich",
			"use_shell": true,
			"command_line": "NAUTILUS=1 ~/maki-immich \"%U\"",
			"min_items": 1,
			"filetypes": ["file"]
		}
	]
}"

