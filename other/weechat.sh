#!/bin/bash

while [ 1 ]; do
	ssh mihari -t ~/quadlets/weechat/attach.sh
	sleep 1
done
