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

func apache2ConfigSaveStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-apache2-config-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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

func sshConfigSaveStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-ssh-config-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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

func mysqlConfigSaveStandalone(filePath string) {
	time := time.Now()
	reportPath := filePath + "sherlock-mysql-config-scan-report-" + time.Format("01-02-2006") + "/"
	runCommand("mkdir " + reportPath)

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