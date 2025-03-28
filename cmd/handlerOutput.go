package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func handleOutput(output string, outputAction outputAction, commandError error) {
	output = strings.Trim(output, "\n")
	switch outputAction.redirectionType {
	case nothing:
		handleStdout(output)
	case redirectStdOut:
		handleRedirectStdOut(output, outputAction.FileName, commandError)
	case appendStdOut:
		handleAppendStdOut(output, outputAction.FileName, commandError)
	case redirectStdErr:
		handleRedirectStdErr(output, outputAction.FileName, commandError)
	case appendStdErr:
		handleAppendStdErr(output, outputAction.FileName, commandError)
	}
}

func handleAppendStdErr(output, fileName string, commandError error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	errorMessage := ""
	if commandError != nil {
		errorMessage = commandError.Error()
	}
	file.WriteString(errorMessage)

	if output != "" {
		fmt.Println(output)
	}
}

func handleRedirectStdErr(output, fileName string, commandError error) {
	errorMessage := ""
	if commandError != nil {
		errorMessage = commandError.Error()
	}
	err := os.WriteFile(fileName, []byte(errorMessage), 0644)
	if err != nil {
		log.Println(err)
	}
	if output != "" {
		fmt.Println(output)
	}
}

func handleAppendStdOut(output, fileName string, commandError error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	if len(output) > 0 {
		output = output + "\n"
	}
	file.WriteString(output)

	errorMessage := ""
	if commandError != nil {
		errorMessage = commandError.Error()
		errorMessage = strings.TrimSuffix(errorMessage, "\n")
	}
	if commandError != nil && len(errorMessage) > 0 {
		fmt.Println(errorMessage)
	}
}

func handleRedirectStdOut(output, fileName string, commandError error) {
	err := os.WriteFile(fileName, []byte(output+"\n"), 0644)
	if err != nil {
		log.Println(err)
	}

	if commandError != nil && len(commandError.Error()) > 0 {
		fmt.Println(strings.Trim(commandError.Error(), "\n"))
	}
}

func handleStdout(output string) {
	output = strings.Trim(output, "\n")
	if output != "" {
		fmt.Println(output)
	}
}
