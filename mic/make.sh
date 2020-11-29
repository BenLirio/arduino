#!/bin/bash

PORT=/dev/cu.usbmodem1421

clear
arduino-cli compile
arduino-cli upload -p "${PORT}"
cat "${PORT}"
