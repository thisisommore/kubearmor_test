package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		time.Sleep(10 * time.Second)
		executeCommands()
	}
}

func executeCommands() {
	catCmd := exec.Command("cat", "/etc/os-release")
	catOutput, err := catCmd.Output()
	if err != nil {
		fmt.Printf("Failed to execute cat command: %v\n", err)
		return
	}

	echoCmd := exec.Command("echo", "Hello, Golang!")
	var echoOutput bytes.Buffer
	echoCmd.Stdout = &echoOutput
	err = echoCmd.Run()
	if err != nil {
		fmt.Printf("Failed to execute echo command: %v\n", err)
		return
	}

	result := fmt.Sprintf("cat output:\n%s\necho output:\n%s", strings.TrimSpace(string(catOutput)), strings.TrimSpace(echoOutput.String()))
	fmt.Println(result)
}
