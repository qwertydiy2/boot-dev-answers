package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) > 4 && len(password) < 13 {
		hasUppercase := false // Flag to track if an uppercase letter is found
		hasDigit := false

		// Iterate over the string using a range loop
		for _, r := range password {
			// Check if the current rune is an uppercase letter
			if unicode.IsUpper(r) {
				hasUppercase = true
			}
			if unicode.IsDigit(r) {
				hasDigit = true
			}
		}
		if hasUppercase && hasDigit {
			return true
		}
	}
	return false
}

type testCase struct {
	password string
	isValid  bool
}

func main() {
	withSubmit := true // Change as needed

	runCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
	}

	submitCases := append(runCases, []testCase{
		{"VeryLongPassword1", false},
		{"Short", false},
		{"1234short", false},
		{"Short5", true},
		{"P4ssword", true},
		{"AA0Z9", true},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	// Simulate *testing.T for main, just print errors
	type fakeT struct{}
	var t fakeT
	tRun := func(name string, f func(t *fakeT)) { f(&t) }
	tErrorf := func(format string, args ...interface{}) {
		fmt.Printf(format, args...)
	}

	for i, test := range testCases {
		tRun(fmt.Sprintf("TestCase%d", i+1), func(t *fakeT) {
			result := isValidPassword(test.password)
			if result != test.isValid {
				failCount++
				tErrorf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Fail
`, test.password, test.isValid, result)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Pass
`, test.password, test.isValid, result)
			}
		})
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}
