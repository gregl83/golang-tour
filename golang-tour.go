package main

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
)

func run(module, lesson string) {
	pathTemplate := "src/github.com/gregl83/golang-tour/%s/%s.go"
	commandPath := fmt.Sprintf(pathTemplate, module, lesson)
	command := exec.Command("/usr/bin/go", "run", commandPath)

	fmt.Printf("Running command: %s", commandPath)

	var stdOut, stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	command.Run()

	fmt.Println()
	fmt.Println()

	fmt.Println(stdOut.String())

	if stdError := stdErr.String(); stdError != "" {
		fmt.Println(stdError)
		fmt.Println("Command failed")
	} else {
		fmt.Println("Command completed")
	}
}

func main() {
	module := os.Args[1]
	lesson := os.Args[2]

	run(module, lesson)
}