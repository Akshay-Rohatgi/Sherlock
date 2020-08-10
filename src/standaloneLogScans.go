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

func getNginxLogsStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-nginx-log-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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