package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Executor ...
func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	} else if s == "bash" {
		cmd := exec.Command("bash")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	cmd := exec.Command("/bin/sh", "-c", "teamwork-go "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}

	return
}
