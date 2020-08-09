package main

import (
	"os"
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
	// Attempt to find apache log files
	getApache2Logs(reportPath)
	// Attempt to find nginx log files
	getNginxLogs(reportPath)
	// Attempt to get apache2 config
	apache2ConfigSave(reportPath)
	// Attempt to get nginx config
	nginxConfigSave(reportPath)
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

func getApache2Logs(filePath string) {
	reportPath := filePath

	// check to see where the access logs are and if found then copy them to desired file path
	if _, err := os.Stat("/var/log/apache/access.log"); err == nil {
		fmt.Println("The apache access log exists at /var/log/apache/access.log!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /var/log/apache/access.log " + reportPath)
	} else if _, err := os.Stat("/var/log/apache2/access.log"); err == nil {
		fmt.Println("The apache access log exists at /var/log/apache2/access.log!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /var/log/apache2/access.log " + reportPath)
	} else if _, err := os.Stat("/etc/httpd/logs/access_log"); err == nil {
		fmt.Println("The apache access log exists at /etc/httpd/logs/access_log!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /etc/httpd/logs/access_log " + reportPath)
	} else {
		fmt.Println("Could not find apache logs!")
	}
}

func getNginxLogs(filePath string) {
	reportPath := filePath

	// checks for nginx access and error logs

	if _, err := os.Stat("/var/log/nginx/access.log"); err == nil {
		fmt.Println("The nginx access log exists at /var/log/nginx/access.log!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /var/log/nginx/access.log " + reportPath)
	} else {
		fmt.Println("Could not find nginx access logs!")
	}


	if _, err := os.Stat("/var/log/nginx/error.log"); err == nil {
		fmt.Println("The nginx access log exists at /var/log/nginx/error.log!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /var/log/nginx/error.log " + reportPath)
	} else {
		fmt.Println("Could not find nginx error logs!")
	}

}

func nginxConfigSave(filePath string) {
	reportPath := filePath

	if _, err := os.Stat("/etc/nginx/nginx.conf"); err == nil {
		fmt.Println("The nginx main configuration file exists at /etc/nginx/nginx.conf!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /etc/nginx/nginx.conf " + reportPath)
	} else {
		fmt.Println("Error")
	}

	if _, err := os.Stat("/etc/nginx/sites-available"); err == nil {
		fmt.Println("Saving all available nginx sites")
		runCommand("cp /etc/nginx/sites-available " + reportPath)
	} else {
		fmt.Println("Error")
	}

}

func apache2ConfigSave(filePath string) {
	reportPath := filePath

	if _, err := os.Stat("/etc/apache2/apache2.conf"); err == nil {
		fmt.Println("The apache main configuration file exists at /etc/apache2/apache2.conf!")
		runCommand("chmod 777 " + reportPath)
		runCommand("cp /etc/apache2/apache2.conf " + reportPath)
	} else {
		fmt.Println("Error")
	}

	if _, err := os.Stat("/etc/apache2/sites-available"); err == nil {
		fmt.Println("Saving all available apache2 sites")
		runCommand("cp /etc/apache2/sites-available " + reportPath)
	} else {
		fmt.Println("Error")
	}

}