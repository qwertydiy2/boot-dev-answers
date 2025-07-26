package main

import (
	"fmt"
	"strconv"
)

type formatter interface{ format() string }

type plaintext struct{ text string }

func (p plaintext) format() string { return p.text }

type bold struct{ text string }

func (b bold) format() string { return "**" + b.text + "**" }

type code struct{ text string }

func (c code) format() string { return "`" + c.text + "`" }

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format()
}

func main() {
	type testCase struct {
		format   formatter
		expected string
	}

	runCases := []testCase{
		{plaintext{text: "Hello, World!"}, "Hello, World!"},
		{bold{text: "Bold Message"}, "**Bold Message**"},
		{code{text: "Code Message"}, "`Code Message`"},
	}

	submitCases := append(runCases, []testCase{
		{code{text: ""}, "``"},
		{bold{text: ""}, "****"},
		{plaintext{text: ""}, ""},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		testName := "Test Case " + strconv.Itoa(i+1)
		formattedMessage := sendMessage(test.format)
		if formattedMessage != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, testName, test.format, test.expected, formattedMessage)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, testName, test.format, test.expected, formattedMessage)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

var withSubmit = true