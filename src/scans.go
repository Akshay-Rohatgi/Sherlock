package main

import (
	"fmt"
)

func systemScan(filePath string) {

	fmt.Println("Starting system scan now!")
	fmt.Println("Saving all files to " + filePath)

	// Get iptable rules and write to file
	runCommand("touch " + filePath + "IPv4IptablesRules.txt")
	runCommand("iptables -S > " + filePath + "IPv4IptablesRules.txt")



}

func firewallScan(filePath, ipVersion string) {
	if ipVersion == "ipv4" {
		runCommand("touch " + filePath + "IPv4IptablesRules.txt")
		runCommand("iptables -S > " + filePath + "IPv4IptablesRules.txt")
	} else if ipVersion == "ipv6" {
		runCommand("touch " + filePath + "IPv6IptablesRules.txt")
		runCommand("ip6tables -S > " + filePath + "IPv6IptablesRules.txt")
	}
}