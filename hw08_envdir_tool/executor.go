package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	okCode   int = 0
	failCode int = 1
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return failCode
	}

	cmdResult := exec.Command(cmd[0], cmd[1:]...) //nolint
	cmdResult.Stdout = os.Stdout
	cmdResult.Stderr = os.Stderr
	for key, value := range env {
		cmdResult.Env = append(cmdResult.Env, fmt.Sprintf("%s=%s", key, value))
	}

	if err := cmdResult.Run(); err != nil {
		return failCode
	}

	return okCode
}
