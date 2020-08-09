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
A tool built for blue teams and incident response teams
sherlock@` + hostname + ` is ready
Use ./src help for a list of commands				  `)

	switch arg := nicerArgs[0]; arg {
	case "help":
		argText := `
========================================================================================================================
Argument list:
help - display possible arguments (what you see right now)
systemScan /path/to/dir - Will save all system log files, iptable rules, and suspicious files to the specified directory
========================================================================================================================`
		fmt.Println(argText)
	default:
		fmt.Println("No arguments specified, exiting with a status code of 0")
		os.Exit(0)
	}

}