package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint
	command.Stdout = os.Stdout

	for k, v := range env {
		if !v.NeedRemove {
			os.Setenv(k, v.Value)
		} else {
			os.Unsetenv(k)
		}
	}

	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		log.Fatalf("Wait run: %v", err)
	}

	return
}
