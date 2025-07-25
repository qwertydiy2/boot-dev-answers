package main

import (
	"fmt"
)

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

func main() {
	type testCase struct {
		emp      employee
		expected int
	}

	runCases := []testCase{
		{fullTime{name: "Bob", salary: 7300}, 7300},
		{contractor{name: "Jill", hourlyPay: 872, hoursPerYear: 982}, 856304},
	}

	submitCases := append(runCases, []testCase{
		{fullTime{name: "Alice", salary: 10000}, 10000},
		{contractor{name: "John", hourlyPay: 1000, hoursPerYear: 1000}, 1000000},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		salary := test.emp.getSalary()
		if salary != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Fail
`, test.emp, test.expected, salary)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  %v
Actual:     %v
Pass
`, test.emp, test.expected, salary)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}
