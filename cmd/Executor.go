package cmd

import (
	"os/exec"
	"os"
	"strings"
	"fmt"
	"log"
)

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	} else if s == "log" {
		cmd := exec.Command("/bin/sh", "-c", "teamwork-go "+s)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	cmd := exec.Command("/bin/sh", "-c", "teamwork-go "+s)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	return
}