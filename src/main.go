package main

import (
	// "flag"
	"fmt"
	"os"
	"os/exec"
	// "strings"
	// "runtime"
)

func main() {

	// get args
	args := os.Args
	nicerArgs := args[1:]
	fmt.Println(nicerArgs)

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
													  `)
	fmt.Println(commandOutput("ls"))
}

func commandOutput(command string) string {
	out, err := exec.Command(command).Output()
	if err != nil {
		fmt.Println(`[ERROR] Error running ` + command)
	}
	return string(out[:])
}