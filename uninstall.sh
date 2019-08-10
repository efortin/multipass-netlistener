#!/bin/bash

echo "Deactivate the Daemon"
launchctl unload -w /Library/LaunchDaemons/networklistener.plist

echo "Remove daemon"
rm /Library/LaunchDaemons/networklistener.plist

echo "Remove application"
rm /usr/local/bin/networklistener

echo "It finish, networklistener is uninstalled."
