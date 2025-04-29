package main

import (
	"fmt"
	"testing"
	"time"
)

// Main Code
func sendMessage(msg message) (string, int) {
	returnedMessage := msg.getMessage()
	return returnedMessage, len(returnedMessage) * 3
}

type message interface {
	getMessage() string
}

// Define birthdayMessage
type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

// Define sendingReport
type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

// Main function
func main() {
	// Example usage
	birthday := birthdayMessage{
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
		recipientName: "John Doe",
	}

	report := sendingReport{
		reportName:    "Monthly Report",
		numberOfSends: 42,
	}

	birthdayMsg, birthdayCost := sendMessage(birthday)
	fmt.Printf("Message: %s\nCost: %d\n\n", birthdayMsg, birthdayCost)

	reportMsg, reportCost := sendMessage(report)
	fmt.Printf("Message: %s\nCost: %d\n", reportMsg, reportCost)
}

// Unit Tests
func TestSendMessage(t *testing.T) {
	type testCase struct {
		msg          message
		expectedText string
		expectedCost int
	}

	testCases := []testCase{
		{
			msg: birthdayMessage{
				birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
				recipientName: "John Doe",
			},
			expectedText: "Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			expectedCost: 168,
		},
		{
			msg: sendingReport{
				reportName:    "First Report",
				numberOfSends: 10,
			},
			expectedText: `Your "First Report" report is ready. You've sent 10 messages.`,
			expectedCost: 183,
		},
		{
			msg: birthdayMessage{
				birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
				recipientName: "Bill Deer",
			},
			expectedText: "Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z",
			expectedCost: 171,
		},
		{
			msg: sendingReport{
				reportName:    "Second Report",
				numberOfSends: 20,
			},
			expectedText: `Your "Second Report" report is ready. You've sent 20 messages.`,
			expectedCost: 186,
		},
	}

	for _, test := range testCases {
		text, cost := sendMessage(test.msg)
		if text != test.expectedText || cost != test.expectedCost {
			t.Errorf("Failed:\nMessage: %+v\nExpected Text: %v, Cost: %v\nGot Text: %v, Cost: %v\n",
				test.msg, test.expectedText, test.expectedCost, text, cost)
		}
	}
}
