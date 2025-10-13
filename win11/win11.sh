#!/bin/bash

if [[ $1 == "stop" ]]; then
	sudo virsh shutdown win11
	exit 0
fi

if [[ $1 == "attach" ]]; then
	vendor_id=$2
	product_id=$3

	# get ids from lsusb
	if [[ $2 == "procontroller" ]]; then
		# nintendo switch pro controller over usb
		vendor_id="057e" 
		product_id="2009" 
	elif [[ $2 == "8bitdo" ]]; then
		# 8bitdo pro 2
		vendor_id="2dc8" 
		product_id="3106" 
	elif [[ $2 == "webcam" ]]; then
		# logitech c920
		vendor_id="046d" 
		product_id="082d" 
	fi

	if [[ $vendor_id == "" && $product_id == "" ]]; then
		echo "Attach: procontroller, 8bitdo, webcam<vendor id> <product id>"
		echo "or manually specify <vendor id> <product id>"
		exit 0
	fi

cat << EOF > /tmp/win11_usb_device.xml
<hostdev mode="subsystem" type="usb" managed="yes">
	<source>
		<vendor id="0x$vendor_id"/>
		<product id="0x$product_id"/>
	</source>
</hostdev>
EOF

	sudo virsh attach-device win11 --live /tmp/win11_usb_device.xml

	rm -f /tmp/win11_usb_device.xml

	exit 0
fi

echo "Starting virtual machine..."
sudo virsh start win11

: '
echo "Starting looking glass..."

sudo chmod 777 /dev/shm/looking-glass

# export __NV_DISABLE_EXPLICIT_SYNC=1

# https://github.com/NVIDIA/egl-wayland/issues/149

# setsid nohup \
looking-glass-client \
-p 5900 -m KEY_RIGHTSHIFT -K 144 -S -F
# 2>/dev/null 1>/dev/null &
'
