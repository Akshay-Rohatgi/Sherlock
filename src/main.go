package main

import (
	// "flag"
	"fmt"
	"os"
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

	fmt.Println(`
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
A tool built by Akshay Rohatgi for incident response teams
sherlock@` + hostname + ` is ready
Use ./src help for a list of commands				  `)

	switch arg := nicerArgs[0]; arg {
	
	case "help":
	
		argText := `

Argument list:


	Neutral args

		help - display possible arguments (what you see right now)
	
	Utility Arguments
		
		localUsers - display all users on the system (including system users)

	System Arguments

		systemScan /path/to/dir - Will save all system log files, service log files, iptable rules, and suspicious files to the specified directory.
		firewallScan /path/to/dir ipv? (ipv4 or ipv6) - Will save ipv4 or ipv6 rules to specified directory. 
		suidScan /path/to/dir - Will save a list of the SUID files on the system to the specified directory.
		sgidScan /path/to/dir - Will save a list of the SGID files on the system to the specified directory.

	Critical Service Arguments:

		servicesSupported - Returns the list of services supported

	Log file Arguments:

		fetch<CRITICAL SERVICE/PROGRAM>Logs /path/to/dir - Fetches log files for that critical service, EX: fetchApache2Logs /home/reports
		fetchAuthLogs

`
		fmt.Println(argText)
	
	case "systemScan":
	
		systemScan(nicerArgs[1])
	
	case "localUsers":

		getLocalUsers()

	case "firewallScan":
	
		firewallScanStandalone(nicerArgs[1], nicerArgs[2])
	
	case "suidScan":
	
		suidScanStandalone(nicerArgs[1])
	
	case "sgidScan":
	
		sgidScanStandalone(nicerArgs[1])
	
	case "servicesSupported":

		supported := `
Apache2
		`
		fmt.Println(supported)

	case "fetchApache2Logs": 

		getApache2LogsStandalone(nicerArgs[1])
	case "fetchAuthLogs":

		getAuthLogsStandalone(nicerArgs[1])

	default:
		fmt.Println("No arguments specified, exiting with a status code of 0")
		os.Exit(0)
	}

}