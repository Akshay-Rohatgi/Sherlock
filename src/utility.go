package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"os/exec"
	"strings"
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

func getLocalUsers() {
	fle, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatal(err.Error())
	}
	scanner := bufio.NewScanner(fle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		user := strings.Split(scanner.Text(), ":")
		fmt.Println(user[0])
	}
	fle.Close()
}