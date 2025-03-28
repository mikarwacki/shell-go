package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var commandMap = make(map[string]func([]string) (string, error))
var userHomeDir string

func main() {
	commandMap["exit"] = handleExit
	commandMap["echo"] = handleEcho
	commandMap["type"] = handleType
	commandMap["pwd"] = handlePwd
	commandMap["cd"] = handleCd
	userHomeDir = os.Getenv("HOME")

	for {
		fmt.Fprint(os.Stdout, "$ ")

		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		clearInput := userInput[:len(userInput)-1]
		cmd, params, outputAction := splitByQuotes(clearInput)
		var output string
		if handleFunc, ok := commandMap[cmd]; ok {
			output, err = handleFunc(params)
		} else {
			output, err = handleInvalidCommand(cmd, params)
		}
		handleOutput(output, outputAction, err)
	}
}
