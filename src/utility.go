package main

import (
	"fmt"
	"os/exec"
)

func commandOutput(command string) string {
	out, err := exec.Command(command).Output()
	if err != nil {
		fmt.Println(`[ERROR] Error running ` + command)
	}
	return string(out[:])
}