package main

import (
	"fmt"
	"time"
)

func systemScan(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-system-scan-report-" + time.Format("01-02-2006") + "/"

	fmt.Println("Starting system scan now!")

	runCommand("mkdir " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	// Get iptable rules and write to file
	firewallScan(reportPath, "ipv4")
}

func firewallScan(filePath, ipVersion string) {
	reportPath := filePath

	fmt.Println("Starting firewall scan now!")
	runCommand("mkdir " + reportPath)
	runCommand("chmod 777 " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	if ipVersion == "ipv4" {
		runCommand("touch " + reportPath + "IPv4IptablesRules.txt")
		runCommand("iptables -S > " + reportPath + "IPv4IptablesRules.txt")
	} else if ipVersion == "ipv6" {
		runCommand("touch " + reportPath + "IPv6IptablesRules.txt")
		runCommand("ip6tables -S > " + reportPath + "IPv6IptablesRules.txt")
	}
}

