package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"os/exec"
	"strings"
	"crypto/md5"
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
		fmt.Println(err.Error())
	}
	defer fle.Close()

	scanner := bufio.NewScanner(fle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		user := strings.Split(scanner.Text(), ":")
		fmt.Println(user[0])
	}

}

func getMD5(filePath string) []byte {
	fle, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer fle.Close()

	hash := md5.New()
	_, err = io.Copy(hash, fle)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("MD5 Hash for %s is %x\n", filePath, hash.Sum(nil))
	return hash.Sum(nil)
}


func md5Baselines() {
	paths := []string{"/etc/passwd", "/etc/shadow", "/etc/group", "/etc/login.defs", "/etc/shells", "/bin/su", "/etc/hosts.allow", "/etc/hosts.deny", "/etc/hosts", "/etc/fstab"}
	for _, file := range paths {
		getMD5(file)
	}
}

func checkFileExist(filePath string) (bool) {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} 
	return false
}