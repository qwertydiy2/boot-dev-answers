package main

import (
	"fmt"
)

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	filteredMessages := []Message{}
	for _, message := range messages {
		if message.Type() == filterType {
			filteredMessages = append(filteredMessages, message)
		}
	}
	return filteredMessages
}

func main() {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}
	type testCase struct {
		filterType    string
		expectedCount int
		expectedType  string
	}

	runCases := []testCase{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
	}

	submitCases := append(runCases, []testCase{
		{"media", 2, "media"},
		{"text", 2, "text"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		filtered := filterMessages(messages, test.filterType)
		if len(filtered) != test.expectedCount {
			failCount++
			fmt.Printf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Fail
`, i+1, test.filterType, test.expectedCount, len(filtered))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Case %d - Filtering for %s
Expecting:  %d messages
Actual:     %d messages
Pass
`, i+1, test.filterType, test.expectedCount, len(filtered))
		}

		for _, m := range filtered {
			if m.Type() != test.expectedType {
				failCount++
				fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Fail
`, i+1, test.expectedType, m.Type())
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  %s message
Actual:     %s message
Pass
`, i+1, test.expectedType, m.Type())
			}
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
