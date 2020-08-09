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
	// Get suid files and write to file
	suidScan(reportPath)
	// Get sgid files and write to file
	sgidScan(reportPath)
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

func suidScan(filePath string) {
	reportPath := filePath

	fmt.Println("Starting SUID scan now!")
	runCommand("mkdir " + reportPath)
	runCommand("chmod 777 " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	runCommand("touch " + reportPath + "SUIDfiles.txt")
	runCommand("sudo find / -perm /4000 2>/dev/null > " + reportPath + "SUIDfiles.txt")
}

func sgidScan(filePath string) {
	reportPath := filePath

	fmt.Println("Starting SGID scan now!")
	runCommand("mkdir " + reportPath)
	runCommand("chmod 777 " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	runCommand("touch " + reportPath + "SGIDfiles.txt")
	runCommand("sudo find / -perm /2000 2>/dev/null > " + reportPath + "SGIDfiles.txt")
}

