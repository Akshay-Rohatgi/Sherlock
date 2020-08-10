package main

import (
	// "flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	// "os/exec"
	// "strings"
	// "runtime"
)

func main() {

	// get args
	args := os.Args
	nicerArgs := args[1:]

	// grab hostname lol
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "linux"
	}

	color.Red(`
  ██████  ██░ ██ ▓█████  ██▀███   ██▓     ▒█████   ▄████▄   ██ ▄█▀
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██ ▒ ██▒▓██▒    ▒██▒  ██▒▒██▀ ▀█   ██▄█▒ 
░ ▓██▄   ▒██▀▀██░▒███   ▓██ ░▄█ ▒▒██░    ▒██░  ██▒▒▓█    ▄ ▓███▄░ 
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██▀▀█▄  ▒██░    ▒██   ██░▒▓▓▄ ▄██▒▓██ █▄ 
▒██████▒▒░▓█▒░██▓░▒████▒░██▓ ▒██▒░██████▒░ ████▓▒░▒ ▓███▀ ░▒██▒ █▄
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒▓ ░▒▓░░ ▒░▓  ░░ ▒░▒░▒░ ░ ░▒ ▒  ░▒ ▒▒ ▓▒
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░  ░▒ ░ ▒░░ ░ ▒  ░  ░ ▒ ▒░   ░  ▒   ░ ░▒ ▒░
░  ░  ░   ░  ░░ ░   ░     ░░   ░   ░ ░   ░ ░ ░ ▒  ░        ░ ░░ ░ 
      ░   ░  ░  ░   ░  ░   ░         ░  ░    ░ ░  ░ ░      ░  ░   
												  ░        
												  					`)
	color.Blue(`
A tool built by Akshay Rohatgi for incident response teams
Use ./sherlock help for a list of commands`)

	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s%s%s is ready\n\n", red("sherlock"), green("@"), blue(hostname))

	switch arg := nicerArgs[0]; arg {

	case "help":

		argText := `

Argument list:

	Neutral args

		help - display possible arguments (what you see right now)
	
	Utility Arguments
		
		localUsers - display all users on the system (including system users)
		getMD5 /path/to/dir - get md5 checksum of a file (this arg is a WASTE of time, why would you write a go function for something you can do with one linux command smh UNLESS you dont have md5sum installed. Then yea, use go.)
		md5Baselines - output md5 checksums of some important files
		criticalFileBackups /path/to/dir - backs up critical linux files to a specified directory

	System Arguments

		systemScan /path/to/dir - Will save all system log files, service log files, iptable rules, and suspicious files to the specified directory.
		firewallScan /path/to/dir ipv? (ipv4 or ipv6) - Will save ipv4 or ipv6 rules to specified directory. 
		suidScan /path/to/dir - Will save a list of the SUID files on the system to the specified directory.
		sgidScan /path/to/dir - Will save a list of the SGID files on the system to the specified directory.

	Critical Service Arguments:

		servicesSupported - Returns the list of services supported
		<SERVICE NAME>ConfigSave /path/to/dir - Creates backup of the specified critical service configuration files to specified dir

	Log file Arguments:

		fetch<CRITICAL SERVICE/PROGRAM>Logs /path/to/dir - Fetches log files for that critical service, EX: fetchApache2Logs /home/reports
		fetchAuthLogs /path/to/dir - Fetches auth log files
		fetchDpkgLogs /path/to/dir - Fetches dpkg log files 

`
		fmt.Println(argText)

	case "systemScan":

		systemScan(nicerArgs[1])

	case "localUsers":

		getLocalUsers()

	case "getMD5":

		getMD5(nicerArgs[1])

	case "md5Baselines":

		md5Baselines()

	case "criticalFileBackups":

		criticalSystemFileBackupStandalone(nicerArgs[1])

	case "firewallScan":

		firewallScanStandalone(nicerArgs[1], nicerArgs[2])

	case "suidScan":

		suidScanStandalone(nicerArgs[1])

	case "sgidScan":

		sgidScanStandalone(nicerArgs[1])

	case "servicesSupported":

		supported := `
Services Supported:
==================
Apache2
Nginx
OpenSSH-Server
		`
		fmt.Println(supported)

	case "apache2ConfigSave":

		apache2ConfigSaveStandalone(nicerArgs[1])

	case "nginxConfigSave":

		nginxConfigSaveStandalone(nicerArgs[1])

	case "sshConfigSave":

		sshConfigSaveStandalone(nicerArgs[1])

	case "fetchApache2Logs":

		getApache2LogsStandalone(nicerArgs[1])

	case "fetchNginxLogs":

		getNginxLogsStandalone(nicerArgs[1])

	case "fetchAuthLogs":

		getAuthLogsStandalone(nicerArgs[1])

	default:
		fmt.Println("No arguments specified, exiting with a status code of 0")
		os.Exit(0)
	}

}
