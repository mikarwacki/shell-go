package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleExit(params []string) (string, error) {
	num, err := strconv.Atoi(params[0])
	if err != nil || num != 0 {
		return "", fmt.Errorf("Invalid exit code")
	}
	os.Exit(num)
	return "", fmt.Errorf("Error exiting")
}

func handleEcho(params []string) (string, error) {
	return strings.Join(params, " "), nil
}

func handleType(params []string) (string, error) {
	searchCmd := params[0]
	if _, ok := commandMap[searchCmd]; ok {
		return fmt.Sprintf("%s is a shell builtin", searchCmd), nil
	}
	path := os.Getenv("PATH")
	paths := strings.Split(path, ":")
	for _, cmdPath := range paths {
		entries, _ := os.ReadDir(cmdPath)
		for _, entry := range entries {
			if searchCmd == entry.Name() {
				return fmt.Sprintf("%s is %s", searchCmd, fmt.Sprintf("%s/%s", cmdPath, searchCmd)), nil
			}
		}
	}
	result := fmt.Sprintf("%s: not found", searchCmd)
	return result, fmt.Errorf(result)
}

func handlePwd(params []string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Error getting current directory")
	}
	return cwd, nil
}

func handleCd(params []string) (string, error) {
	if params[0] == "~" {
		params = []string{userHomeDir}
	}
	if err := os.Chdir(params[0]); err != nil {
		return fmt.Sprintf("cd: %s: No such file or directory", params[0]), nil
	}
	return "", nil
}
