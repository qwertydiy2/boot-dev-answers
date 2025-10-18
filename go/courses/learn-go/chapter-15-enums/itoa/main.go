package main

import (
	"fmt"
)

type emailStatus int

const (
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return fmt.Errorf("error updating user status: %w", err)
	}
	err = a.track(em.status)
	if err != nil {
		return fmt.Errorf("error tracking user bounce: %w", err)
	}
	return nil
}

func main() {
	type testCase struct {
		status   emailStatus
		expected string
	}

	runCases := []testCase{
		{emailBounced, "emailBounced"},
		{emailInvalid, "emailInvalid"},
		{emailDelivered, "emailDelivered"},
	}

	submitCases := append(runCases, []testCase{
		{emailOpened, "emailOpened"},
		{17, "Unknown"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getEmailStatusName(test.status)
		if output != test.expected {
			failCount++
			fmt.Print(`---------------------------------
Test Failed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  status:   %v
  expected: %v
  actual:   %v
`, test.status, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}

}

func getEmailStatusName(status emailStatus) string {
	switch status {
	case emailBounced:
		return "emailBounced"
	case emailInvalid:
		return "emailInvalid"
	case emailDelivered:
		return "emailDelivered"
	case emailOpened:
		return "emailOpened"
	default:
		return "Unknown"
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
