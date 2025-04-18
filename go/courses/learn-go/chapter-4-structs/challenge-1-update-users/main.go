package main

import (
	"fmt"
	"testing"
)

type membership struct {
    Type             string
    MessageCharLimit int
}

type User struct {
	membership
	Name             string
}

func newUser(name string, membershipType string) User {
    if membershipType == "premium" {
        return User{Name: name, membership: membership{Type: membershipType, MessageCharLimit: 1000}}
    }
    return User{Name: name, membership: membership{Type: membershipType, MessageCharLimit: 100}}
}

func main() {
	// Run tests
	runTests()
}

func runTests() {
	fmt.Println("\nRunning tests...\n")

	// Create a test suite
	var t testing.T

	// Run the test cases
	Test(&t)

	if t.Failed() {
		fmt.Println("\nSome tests failed. See output above.")
	} else {
		fmt.Println("\nAll tests passed!")
	}
}

func Test(t *testing.T) {
	type testCase struct {
		name           string
		membershipType string
	}

	runCases := []testCase{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
	}

	submitCases := append(runCases, []testCase{
		{"Renarin", "standard"},
		{"Lift", "premium"},
		{"Dalinar", "standard"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		user := newUser(test.name, test.membershipType)

		msgCharLimit := 100
		if test.membershipType == "premium" {
			msgCharLimit = 1000
		}

		if user.Name != test.name {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (name):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.name, user.Name)
		} else if user.Type != test.membershipType {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (membership type):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, test.membershipType, user.Type)
		} else if user.MessageCharLimit != msgCharLimit {
			failCount++
			t.Errorf(`---------------------------------
Test Failed (message character limit):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, test.name, test.membershipType, msgCharLimit, user.MessageCharLimit)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v, %v, %v
Actual:     %v, %v, %v
`, test.name, test.membershipType, test.name, test.membershipType, msgCharLimit, user.Name, user.Type, user.MessageCharLimit)
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
