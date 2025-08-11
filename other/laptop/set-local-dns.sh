#!/bin/bash

conf_path=/etc/systemd/resolved.conf.d/dns_servers.conf

# grep \#DNS $conf_path > /dev/null 2> /dev/null

# if [[ $? -eq 0 ]]; then
if [[ $1 == 1 ]]; then
	echo "enabling local dns"
	echo ""
	cat << EOF | sudo tee $conf_path > /dev/null
[Resolve]
DNS=127.0.0.1
Domains=~.
FallbackDNS=
EOF
# else
elif [[ $1 == 0 ]]; then
	echo "disabling local dns"
	echo ""
	cat << EOF | sudo tee $conf_path > /dev/null
[Resolve]
#DNS=127.0.0.1
#Domains=~.
#FallbackDNS=
EOF
else
	echo "usage: 0 to disable, 1 to enable"
	exit 1
fi

sudo systemctl restart systemd-resolved.service

output=$(resolvectl)
line_number=$(echo "$output" | grep -n -m 1 -P "^$" | cut -d: -f1)

cat $conf_path
echo ""

echo "$output" | head -n $(expr $line_number - 1)
