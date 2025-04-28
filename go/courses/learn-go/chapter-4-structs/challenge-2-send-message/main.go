package main

import (
	"fmt"
	"testing"
)

// Corrected method signature and implementation
func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if u.MessageCharLimit > messageLength {
		return message, true
	}
	return "", false
}

// User struct definition
type User struct {
	Name string
	Membership
}

// Membership struct definition
type Membership struct {
	Type             string
	MessageCharLimit int
}

// Constructor for User
func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

// Test cases
func Test(t *testing.T) {
	type testCase struct {
		name           string
		membershipType string
		message        string
		expectResult   string
		expectSuccess  bool
	}

	runCases := []testCase{
		{"Syl", "standard", "Hello, Kaladin!", "Hello, Kaladin!", true},
		{"Pattern", "premium", "You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.", "You are not as good with patterns... You are abstract. You think in lies and tell them to yourselves. That is fascinating, but it is not good for patterns.", true},
		{"Dalinar", "standard", "I will take responsibility for what I have done. If I must fall, I will rise each time a better man.", "I will take responsibility for what I have done. If I must fall, I will rise each time a better man.", true},
	}

	submitCases := append(runCases, []testCase{
		{"Pattern", "standard", "Humans can see the world as it is not. It is why your lies can be so strong. You are able to not admit that they are lies.", "", false},
		{"Dabbid", "premium", ".........................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................................", "", false},
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
		result, pass := user.SendMessage(test.message, len(test.message))
		if test.expectSuccess != pass || result != test.expectResult {
			failCount++
			t.Errorf(`---------------------------------
Test Failed:
* user:               %s
* membership type:    %s
* message:            %s
* expected result:    %s
* expected success:   %v
* actual result:      %s
* actual success:     %v
`, test.name, test.membershipType, test.message, test.expectResult, test.expectSuccess, result, pass)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
* user:               %s
* membership type:    %s
* message:            %s
* expected result:    %s
* expected success:   %v
* actual result:      %s
* actual success:     %v
`, test.name, test.membershipType, test.message, test.expectResult, test.expectSuccess, result, pass)
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

// Main function
func main() {
	// Create a new user
	user := newUser("Kaladin", "premium")

	// Send a message
	message := "Honor is dead. But I'll see what I can do."
	messageLength := len(message)
	result, success := user.SendMessage(message, messageLength)

	// Print the result
	if success {
		fmt.Printf("Message sent successfully: %s\n", result)
	} else {
		fmt.Println("Failed to send message: Message exceeds character limit.")
	}
}
