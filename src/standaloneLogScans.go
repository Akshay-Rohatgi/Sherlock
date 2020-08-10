package main 

import (
	"os"
	"fmt"
	"time"
)

func getApache2LogsStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-apache2-log-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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

func getNginxLogsStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-nginx-log-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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

func getAuthLogsStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-auth-log-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

	if _, err := os.Stat("/var/log/auth.log"); err == nil {
		fmt.Println("The system authentication log exists!")
		runCommand("cp /var/log/auth.log " + reportPath)
	} else {
		fmt.Println("The system authentication log does NOT exist!")
		fmt.Println("sus :thinking:")
	}
}

func getDpkgLogsStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-dpkg-log-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

	if _, err := os.Stat("/var/log/dpkg.log"); err == nil {
		fmt.Println("The dpkg log exists!")
		runCommand("cp /var/log/dpkg.log " + reportPath)
	} else {
		fmt.Println("The dpkg log does NOT exist!")
	}
}