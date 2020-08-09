package main

import (
	"fmt"
	"os/exec"
	// "strings"
)

func commandOutput(command string) string {
	// commandArray := strings.Split(command, " ")
	// fmt.Println(commandArray)
	// for _, component := range commandArray {
	// 	component = `"` + component + `",` 
	// 	fmt.Println(component)
	// }
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		fmt.Println(`[ERROR] Error running ` + command)
		fmt.Println(err)
	}
	return string(out[:])
}

