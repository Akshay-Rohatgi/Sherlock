package main

import (
	"fmt"
)

func systemScan(filePath string) {

	fmt.Println("Starting system scan now!")
	fmt.Println("Saving all files to " + filePath)

	// Get iptable rules
	iptablesRules := commandOutput("/usr/sbin/iptables -S")
	fmt.Println(iptablesRules)
}
