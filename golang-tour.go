package main

import (
	"fmt"
	"os"
	"os/exec"
	"bytes"
	"log"
)

func run(module, lesson string) {
	pathTemplate := "src/github.com/gregl83/golang-tour/%s/%s.go"
	commandPath := fmt.Sprintf(pathTemplate, module, lesson)
	command := exec.Command("/usr/bin/go", "run", commandPath)

	fmt.Printf("Running command: %s", commandPath)

	var stdOut, stdErr bytes.Buffer
	command.Stdout = &stdOut
	command.Stderr = &stdErr
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println(stdErr.String())
	fmt.Println()
	fmt.Println(stdOut.String())
	fmt.Println()
	fmt.Println("Command completed")
}

func main() {
	module := os.Args[1]
	lesson := os.Args[2]

	run(module, lesson)
}