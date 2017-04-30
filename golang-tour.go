package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"bytes"
	"fmt"
)

var (
	_, base, _, _ = runtime.Caller(0)
	basepath = filepath.Dir(base)
)

func run(module, lesson string) {
	commandPath := fmt.Sprintf("%s/%s/%s.go", basepath, module, lesson)
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
	if len(os.Args) < 3 {
		fmt.Println("Module and lesson arguments are required")
		return
	}

	run(os.Args[1], os.Args[2])
}