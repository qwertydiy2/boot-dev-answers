package main

import (
	"fmt"
)

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

func (g groupMessage) importance() int { return g.priorityLevel }

func (s systemAlert) importance() int { return 100 }

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.priorityLevel
	case systemAlert:
		return v.alertCode, 100
	default:
		return "", 0
	}
}

func main() {
	type testCase struct {
		notification       notification
		expectedID         string
		expectedImportance int
	}

	runCases := []testCase{
		{
			directMessage{senderUsername: "Kaladin", messageContent: "Life before death", priorityLevel: 10, isUrgent: true},
			"Kaladin",
			50,
		},
		{
			groupMessage{groupName: "Bridge 4", messageContent: "Soups ready!", priorityLevel: 2},
			"Bridge 4",
			2,
		},
		{
			systemAlert{alertCode: "ALERT001", messageContent: "THIS IS NOT A TEST HIGH STORM COMING SOON"},
			"ALERT001",
			100,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			directMessage{senderUsername: "Shallan", messageContent: "I am that I am.", priorityLevel: 5, isUrgent: false},
			"Shallan",
			5,
		},
		{
			groupMessage{groupName: "Knights Radiant", messageContent: "For the greater good.", priorityLevel: 10},
			"Knights Radiant",
			10,
		},
		{
			directMessage{senderUsername: "Adolin", messageContent: "Duels are my favoritefavourite.", priorityLevel: 3, isUrgent: true},
			"Adolin",
			50,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for i, test := range testCases {
		id, importance := processNotification(test.notification)
		if id != test.expectedID || importance != test.expectedImportance {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Fail
`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed: TestProcessNotification_%d
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Pass
`, i+1, test.notification, test.expectedID, test.expectedImportance, id, importance)
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
