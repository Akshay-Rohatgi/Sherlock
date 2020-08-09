package main

import (
	"fmt"
)

func systemScan(filePath string) {

	fmt.Println("Starting system scan now!")
	fmt.Println("Saving all files to " + filePath)

	// Get iptable rules and write to file
	runCommand("touch " + filePath + "IPv4iptablesRules.txt")
	runCommand("iptables -S > " + filePath + "IPv4IptablesRules.txt")


}
