package main

import (
	"fmt"
	"os"
	"strings"
)

func errorAndExit(err error, exitCode int) {
	fmt.Println(err)
	os.Exit(exitCode)
}

func nthOccurrenceIndex(str string, pattern byte, n int) int {
	if n <= 0 {
		return -1
	}

	i := 0
	for ; i < len(str) && n > 0; i++ {
		if str[i] == pattern {
			n--
		}
	}

	if n > 0 {
		return -1
	}
	return i
}

func main() {
	if len(os.Args) <= 1 {
		errorAndExit(fmt.Errorf("Expected 1 input, received %d instead", len(os.Args)-1), 1)
	}

	inputString := os.Args[1]

	hasDash := strings.Contains(inputString, "-")

	if !hasDash {
		errorAndExit(fmt.Errorf("Invalid input. Expects at least 1 dash (-) in the input"), 1)
	}

	// Input can come in two flavors:
	// 1. kebab-case
	// 2. Sentence case
	// Both input should (must) has JIRA ticket prefixed

	spaceIndex := nthOccurrenceIndex(inputString, ' ', 1)
	if spaceIndex != -1 {
		inputString = inputString[:spaceIndex-1]
	}

	dashIndex := nthOccurrenceIndex(inputString, '-', 2)
	if dashIndex != -1 {
		inputString = inputString[:dashIndex-1]
	}

	fmt.Println(inputString)
}
