package main

import (
	"strings"
)

type RedirectionType int

const (
	redirectStdOut RedirectionType = iota
	redirectStdErr
	appendStdOut
	appendStdErr
	nothing
)

type outputAction struct {
	FileName        string
	redirectionType RedirectionType
}

func splitByQuotes(params string) (string, []string, outputAction) {
	var currWord []byte
	var result []string
	var isBackslash bool

	isInQuotes := map[byte]bool{'"': false, '\'': false}
	validEscapeChars := map[byte]bool{'$': true, '\\': true, '"': true}

	for i := 0; i < len(params); i++ {
		char := params[i]

		if i < len(params)-1 && !isInQuotes['"'] && !isInQuotes['\''] {
			if char == '>' && params[i+1] == '>' {
				fileName := strings.TrimSpace(params[i+2:])
				return result[0], result[1:], outputAction{FileName: fileName, redirectionType: appendStdOut}
			}
			if char == '1' && params[i+1] == '>' && params[i+2] == '>' {
				fileName := strings.TrimSpace(params[i+3:])
				return result[0], result[1:], outputAction{FileName: fileName, redirectionType: appendStdOut}
			}
			if char == '>' || char == '1' && params[i+1] == '>' {
				fileName := strings.TrimSpace(params[i+2:])
				return result[0], result[1:], outputAction{FileName: fileName, redirectionType: redirectStdOut}
			}
		}
		if i < len(params)-1 && !isInQuotes['"'] && !isInQuotes['\''] && (char == '2' && params[i+1] == '>') {
			if params[i+2] == '>' {
				fileName := strings.TrimSpace(params[i+3:])
				return result[0], result[1:], outputAction{FileName: fileName, redirectionType: appendStdErr}
			}
			fileName := strings.TrimSpace(params[i+2:])
			return result[0], result[1:], outputAction{FileName: fileName, redirectionType: redirectStdErr}
		}
		if !isInQuotes['\''] && char == '\\' && !isBackslash {
			isBackslash = true
			continue
		}
		if isBackslash {
			if !validEscapeChars[char] && isInQuotes['"'] {
				currWord = append(currWord, '\\')
			}
			currWord = append(currWord, char)
			isBackslash = false
			continue
		}
		if enclosed, ok := isInQuotes[char]; ok && !areOtherQuotesOpened(isInQuotes, char) {
			if enclosed {
				if i < len(params)-1 && params[i+1] != ' ' {
					if _, ok := isInQuotes[params[i+1]]; ok {
						i++
					}
					continue
				}
				result = append(result, string(currWord))
				currWord = []byte{}
			}
			isInQuotes[char] = !isInQuotes[char]
			continue
		}
		if char == ' ' && !isQuoted(isInQuotes) {
			if len(currWord) > 0 {
				result = append(result, string(currWord))
				currWord = []byte{}
			}
			continue
		}
		currWord = append(currWord, char)
	}

	if len(currWord) > 0 {
		result = append(result, string(currWord))
	}
	return result[0], result[1:], outputAction{FileName: "", redirectionType: nothing}
}

func isQuoted(isInQuotes map[byte]bool) bool {
	for _, v := range isInQuotes {
		if v {
			return true
		}
	}
	return false
}

func areOtherQuotesOpened(quotes map[byte]bool, currentQuote byte) bool {
	for k, v := range quotes {
		if k != currentQuote {
			return v
		}
	}
	return false
}

func splitCommandAndParams(input string) (string, string) {
	if input[0] != '\'' {
		cmd, params, _ := strings.Cut(input, " ")
		return cmd, params
	}
	var cmd []byte
	var params []byte
	quoteReached := false

	for i := 1; i < len(input); i++ {
		if input[i] == '\'' {
			quoteReached = true
			continue
		}
		if quoteReached {
			params = append(params, input[i])
		} else {
			cmd = append(cmd, input[i])
		}
	}
	return string(cmd), string(params)
}
