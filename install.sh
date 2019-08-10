#!/bin/bash

echo "Copy application and fix permissions"
curl -ssf -L https://raw.githubusercontent.com/efortin/multipass-netlistener/master/build/networklistener -o /usr/local/bin/networklistener
chown root:wheel /usr/local/bin/networklistener
chmod 500 /usr/local/bin/networklistener

echo "Copy Daemon config and fix permissions"
curl -ssf -L https://raw.githubusercontent.com/efortin/multipass-netlistener/master/deployments/networklistener.plist -o /Library/LaunchDaemons/networklistener.plist
chown root:wheel /Library/LaunchDaemons/networklistener.plist
chmod 644 /Library/LaunchDaemons/networklistener.plist

echo "Activate the Daemon"
sudo launchctl load -w /Library/LaunchDaemons/networklistener.plist
sudo launchctl start /Library/LaunchDaemons/networklistener.plist

echo "Great, networklistener is ready to serve ..."
