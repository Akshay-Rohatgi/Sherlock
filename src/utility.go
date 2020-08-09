package main

import (
	"fmt"
	"os/exec"
	// "strings"
)

func commandOutput(command string) string {
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Println(`[ERROR] Error running ` + command)
		fmt.Println(err)
	}
	return string(out[:])
}

func runCommand(command string) {
	run := exec.Command("sh", "-c", command)
	run.Run()
}