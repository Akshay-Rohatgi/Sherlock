package main 

import (
	// "os"
	"fmt"
	"time"
	"github.com/fatih/color"
)

func firewallScanStandalone(filePath, ipVersion string) {
	time := time.Now()
	reportPath := filePath + "sherlock-firewall-scan-report-" + time.Format("01-02-2006") + "/"

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

func suidScanStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-suid-scan-report-" + time.Format("01-02-2006") + "/"

	fmt.Println("Starting SUID scan now!")
	runCommand("mkdir " + reportPath)
	runCommand("chmod 777 " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	runCommand("touch " + reportPath + "SUIDfiles.txt")
	runCommand("sudo find / -perm /4000 2>/dev/null > " + reportPath + "SUIDfiles.txt")
}

func sgidScanStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-sgid-scan-report-" + time.Format("01-02-2006") + "/"

	fmt.Println("Starting SGID scan now!")
	runCommand("mkdir " + reportPath)
	runCommand("chmod 777 " + reportPath)
	fmt.Println("Saving all files to " + reportPath)

	runCommand("touch " + reportPath + "SGIDfiles.txt")
	runCommand("sudo find / -perm /2000 2>/dev/null > " + reportPath + "SGIDfiles.txt")
}

func criticalSystemFileBackupStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-critical-files-backup" + time.Format("01-02-2006") + "/"
	paths := []string{"/etc/passwd", "/etc/shadow", "/etc/group", "/etc/login.defs", "/etc/shells", "/bin/su", "/etc/hosts.allow", "/etc/hosts.deny", "/etc/hosts", "/etc/fstab"}
	
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()

	runCommand("mkdir " + reportPath)

	for _, file := range paths {

		if checkFileExist(file) == true {
			fmt.Printf("[%s] Saving file to specified path\n", green("FOUND"))
			runCommand("cp " + file + " " + reportPath)
		} else {
			fmt.Printf("[%s] Cannot save the file to specified path", red("NOT FOUND"))
		}
	}
	fmt.Printf("[%s] Cannot save the file to specified path\n", blue("COMPLETE"))
}