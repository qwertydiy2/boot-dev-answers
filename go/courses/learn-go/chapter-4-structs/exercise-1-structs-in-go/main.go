package main

import (
	"fmt"
	"testing"
)

// messageToSend represents a message with a phone number and the message content
type messageToSend struct {
	phoneNumber int
	message     string
}

// getMessageText generates a formatted message string
func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}

func main() {
	// Example usage
	message := messageToSend{
		phoneNumber: 1234567890,
		message:     "Hello, World!",
	}
	fmt.Println(getMessageText(message))

	// Run tests
	runTests()
}

func runTests() {
	fmt.Println("\nRunning tests...\n")

	// Create a test suite
	var t testing.T

	// Run the test cases
	TestGetMessageText(&t)

	if t.Failed() {
		fmt.Println("\nSome tests failed. See output above.")
	} else {
		fmt.Println("\nAll tests passed!")
	}
}

// TestGetMessageText contains the test cases for getMessageText
func TestGetMessageText(t *testing.T) {
	type testCase struct {
		phoneNumber int
		message     string
		expected    string
	}

	runCases := []testCase{
		{148255510981, "Thanks for signing up", "Sending message: 'Thanks for signing up' to: 148255510981"},
		{148255510982, "Love to have you aboard!", "Sending message: 'Love to have you aboard!' to: 148255510982"},
	}

	submitCases := append(runCases, []testCase{
		{148255510983, "We're so excited to have you", "Sending message: 'We're so excited to have you' to: 148255510983"},
		{148255510984, "", "Sending message: '' to: 148255510984"},
		{148255510985, "Hello, World!", "Sending message: 'Hello, World!' to: 148255510985"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMessageText(messageToSend{
			phoneNumber: test.phoneNumber,
			message:     test.message,
		})
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.phoneNumber, test.message, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.phoneNumber, test.message, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
