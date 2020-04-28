package main

import (
	"log"
	"os"
)

const MinArgs = 3

func main() {
	logger := log.New(os.Stderr, "", 0)

	if len(os.Args) < MinArgs {
		logger.Printf("Not enough arguments. Got %d, expected >= %d", len(os.Args), MinArgs)
	}
	envPath, commandWithArgs := os.Args[1], os.Args[2:]

	envVars, err := ReadDir(envPath)
	if err != nil {
		logger.Println(err)
		return
	}

	RunCmd(commandWithArgs, envVars)
}
