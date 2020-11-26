#!/bin/bash

if [[ $(whoami) != "root" ]]
then
	echo "Must execute as root"
	exit 1
fi

BOARD="arduino:avr:mega"
PORT="/dev/ttyACM0"

clear
echo "BOARD=${BOARD}"
echo "PORT=${PORT}"
arduino-cli compile -b "${BOARD}"
arduino-cli upload -b "${BOARD}" -p "${PORT}"
