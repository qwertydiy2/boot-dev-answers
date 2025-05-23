package main

import (
	"fmt"
	"testing"
)

type authenticationInfo struct {
	username string
	password string
}

// getBasicAuth generates a basic authorization string
func (user authenticationInfo) getBasicAuth() string {
	authString := "Authorization: Basic " + user.username + ":" + user.password
	return authString
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
	TestGetBasicAuth(&t)

	if t.Failed() {
		fmt.Println("\nSome tests failed. See output above.")
	} else {
		fmt.Println("\nAll tests passed!")
	}
}

func TestGetBasicAuth(t *testing.T) {
	type testCase struct {
		auth     authenticationInfo
		expected string
	}

	runCases := []testCase{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}

	submitCases := append(runCases, []testCase{
		{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := test.auth.getBasicAuth()
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Fail
`, test.auth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %s
Actual:     %s
Pass
`, test.auth, test.expected, output)
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
