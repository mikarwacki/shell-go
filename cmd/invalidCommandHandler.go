package main

import (
	"fmt"
	"io"
	"os/exec"
)

func handleInvalidCommand(cmd string, params []string) (string, error) {
	_, err := handleType([]string{cmd})
	if err != nil {
		fmt.Printf("%s: not found\n", cmd)
		return "", fmt.Errorf("%s: not found\n", cmd)
	}
	execCmd := exec.Command(cmd, params...)
	stdoutPipe, err := execCmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stdout pipe: %v", err)
	}

	stderrPipe, err := execCmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stderr pipe: %v", err)
	}

	if err := execCmd.Start(); err != nil {
		return "", fmt.Errorf("error starting command: %v", err)
	}
	stdoutBytes, err := io.ReadAll(stdoutPipe)
	if err != nil {
		return "", fmt.Errorf("error reading stdout: %v", err)
	}

	stderrBytes, err := io.ReadAll(stderrPipe)
	if err != nil {
		return "", fmt.Errorf("error reading stderr: %v", err)
	}

	return (string(stdoutBytes)), fmt.Errorf("%s", string(stderrBytes))
}
