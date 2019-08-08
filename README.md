
# Multipass NetListener

Multipass NetListener is a LaunchDaemon that remove hostfilter
on MacOSX default bridge for each vm.

You should be use it, if you add multiple network addresses
inside your multipass vm. By default, Mac osx use L2 Filtering
on bridge (man ifconfig @hostfilter).

Thank to @yadel for network hack !

You can see it with the following command:
```bash
ifconfig -v bridge100

#output ( see hostfilter )
  ...
        member: en3 flags=3<LEARNING,DISCOVER>
                ifmaxaddr 0 port 7 priority 0 path cost 0
                hostfilter 0 hw: 0:0:0:0:0:0 ip: 0.0.0.0
        member: en4 flags=3<LEARNING,DISCOVER>
                ifmaxaddr 0 port 12 priority 0 path cost 0
                hostfilter 0 hw: 0:0:0:0:0:0 ip: 0.0.0.0
  ...
```

## Installation

```bash
# Copy daemon and fix permissions
sudo cp networklistener.plist /Library/LaunchDaemons/
sudo chown root:wheel /Library/LaunchDaemons/networklistener.plist
sudo chmod 644 /Library/LaunchDaemons/networklistener.plist

#Copy script and fix permissions
sudo cp networklistener.sh /usr/local/bin/networklistener
sudo chown root:wheel /usr/local/bin/networklistener
sudo chmod 500 /usr/local/bin/networklistener

#Activate permanently the Daemon
sudo launchctl load -w /Library/LaunchDaemons/networklistener.plist
sudo launchctl start /Library/LaunchDaemons/networklistener.plist

```

## References

- https://medium.com/@fahimhossain_16989/adding-startup-scripts-to-launch-daemon-on-mac-os-x-sierra-10-12-6-7e0318c74de1
- https://apple.stackexchange.com/questions/32354/how-do-you-run-a-script-after-a-network-interface-comes-up/97751
- https://apple.stackexchange.com/questions/350196/triggering-launch-agents-for-path-in-sandboxed-app-using-watchpaths-does-not-wor
