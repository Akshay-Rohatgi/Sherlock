package main 

import (
	"os"
	"fmt"
	"time"
)

func nginxConfigSaveStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-nginx-config-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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

func apache2ConfigSaveStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-apache2-config-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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