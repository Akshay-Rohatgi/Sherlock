package main

import (
	"os"
	"fmt"
	"time"
	"github.com/fatih/color"
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
	// Attempt to get ssh config
	sshConfigSave(reportPath)
	// Attempt to get mysql config
	mysqlConfigSave(reportPath)
	// critical files backup
	criticalSystemFileBackup(reportPath)
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
	possibleLogFileLocations := []string{"/var/log/apache/access.log", "/var/log/apache2/access.log", "/etc/httpd/logs/access_log"}

	// check to see where the access logs are and if found then copy them to desired file path
	for _, file := range possibleLogFileLocations {
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The Apache access log exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The Apache access log does not exist at " + file)
		}
	}
}

func getNginxLogs(filePath string) {
	reportPath := filePath

	// checks for nginx access and error logs
	possibleLogFileLocations := []string{"/var/log/nginx/access.log", "/var/log/nginx/error.log"}

	for _, file := range possibleLogFileLocations {
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The Nginx log exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The Nginx log does not exist at " + file)
		}
	}

}

func nginxConfigSave(filePath string) {
	reportPath := filePath
	possibleFileLocations := []string{"/etc/nginx/nginx.conf", "/etc/nginx/sites-available"}

	for _, file := range possibleFileLocations {
		fmt.Println(file)
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The Nginx configuration file exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The Nginx configuration file does not exist at " + file)
		}
	}

}

func apache2ConfigSave(filePath string) {
	reportPath := filePath
	possibleFileLocations := []string{"/etc/apache2/apache2.conf", "/etc/apache2/sites-available"}

	for _, file := range possibleFileLocations {
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The Apache2 configuration file exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The Apache2 configuration file does not exist at " + file)
		}
	}

}

func criticalSystemFileBackup(filePath string) {
	reportPath := filePath
	paths := []string{"/etc/passwd", "/etc/shadow", "/etc/group", "/etc/login.defs", "/etc/shells", "/bin/su", "/etc/hosts.allow", "/etc/hosts.deny", "/etc/hosts", "/etc/fstab"}
	
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

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


func sshConfigSave(filePath string) {
	reportPath := filePath

	possibleFileLocations := []string{"/etc/ssh/sshd_config"}

	for _, file := range possibleFileLocations {
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The SSH configuration file exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The SSH configuration file does not exist at " + file)
		}
	}
}


func mysqlConfigSave(filePath string) {
	reportPath := filePath

	possibleFileLocations := []string{"/etc/mysql/my.cnf", "/etc/mysql/mysql.conf.d/mysqld.cnf"}

	for _, file := range possibleFileLocations {
		if _, err := os.Stat(file); err == nil {
			fmt.Println("The MySQL configuration file exists at " + file)
			runCommand("chmod 777 " + reportPath)
			runCommand("cp -r " + file + " " + reportPath)
		} else {
			errorPrint("The MySQL configuration file does not exist at " + file)
		}
	}

}