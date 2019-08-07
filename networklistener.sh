#!/bin/bash

echo "Looking interfaces changes..." >> /tmp/networkhost
sleep 10
echo "Applying changes.." >> /tmp/networkhost
ifconfig -v bridge100 | grep "member: en" | awk '{ system("sudo ifconfig bridge100 -hostfilter "$2)}'
