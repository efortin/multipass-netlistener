package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"net"
	"os"
	"os/exec"
	"strings"
)

type MemberState struct {
	InterfaceName string `json: "name"`
	FilterActive  bool   `json: "filter-status"`
	MacFilter     string `json: "filtered-mac"`
	IpFilter      string `json: "filtered-ip"`
}

func main() {

	bridgeName := flag.String("bridge-name", "bridge100", "The bridge name to watch, default is bridge100")
	flag.Parse()

	cmd := exec.Command("ifconfig", "-v", *bridgeName)
	cmdOutput := &bytes.Buffer{}
	cmdOutputErr := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdOutputErr
	err := cmd.Run()
	output := string(cmdOutput.Bytes())
	if err != nil {
		beeep.Alert(
			"Network Listener",
			fmt.Sprintf("No vm are running or %s doesn't not exists.", *bridgeName),
			"")
		os.Stderr.WriteString(string(cmdOutputErr.Bytes()))
		os.Stderr.WriteString(err.Error())
	}
	memberState, _ := parseIfConfig(output)
	removeHostFiltering(memberState, *bridgeName)

}

func removeHostFiltering(hosts *[]MemberState, bridge string) {
	for _, host := range *hosts {
		if host.FilterActive {
			cmd := exec.Command("ifconfig", bridge, "-hostfilter", host.InterfaceName)
			err := cmd.Run()
			if err != nil {
				beeep.Alert(
					"Network Listener",
					"Network Listener must be executed with privileges.", "")
				os.Stderr.WriteString(err.Error())
				return
			}

			hostJson, _ := json.Marshal(host)
			os.Stdout.WriteString("\nchanging :" + string(hostJson))
			message := fmt.Sprintf("The interface %v became unfiltered \n(ip: %v, mac %v)", host.InterfaceName, host.IpFilter, host.MacFilter)
			beeep.Alert("Network Listener", message, "")

		}
	}
}

func parseIfConfig(output string) (*[]MemberState, error) {

	memberStates := []MemberState{}

	splits := strings.Split(output, "\n")
	for index, line := range splits {
		token := strings.TrimLeft(line, "\t")
		if strings.HasPrefix(token, "member") {
			subline2 := strings.TrimSpace(splits[index+2])

			hostfilterstate := strings.Split(strings.TrimLeft(subline2, "hostfilter "), " ")[0] == "1"
			macfilter := strings.Split(strings.Split(subline2, "hw: ")[1], " ")[0]
			ipfilter := strings.Split(strings.Split(subline2, "ip: ")[1], " ")[0]
			ip := net.ParseIP(ipfilter)

			desc := MemberState{
				InterfaceName: strings.Split(strings.TrimLeft(token, "member:"), " ")[1],
				FilterActive:  hostfilterstate,
				MacFilter:     macfilter,
				IpFilter:      ip.String(),
			}
			memberStates = append(memberStates, desc)
		}
	}
	return &memberStates, nil

}
