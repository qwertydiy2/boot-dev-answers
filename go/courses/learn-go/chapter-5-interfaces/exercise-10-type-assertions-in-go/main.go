package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}

	runCases := []testCase{
		{
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}

	submitCases := append(runCases, []testCase{
		{
			email{
				isSubscribed: false,
				body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
				toAddress:    "soldier@monty.com",
			},
			"soldier@monty.com",
			6.95,
		},
		{
			email{
				isSubscribed: true,
				body:         "Pull the other one!",
				toAddress:    "arthur@monty.com",
			},
			"arthur@monty.com",
			0.19,
		},
		{
			sms{
				isSubscribed:  true,
				body:          "I am. And this my trusty servant Patsy.",
				toPhoneNumber: "+155555509832",
			},
			"+155555509832",
			1.17,
		},
		{
			invalid{},
			"",
			0.0,
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	passCount := 0
	failCount := 0
	skipped := len(submitCases) - len(testCases)

	for _, test := range testCases {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			fmt.Errorf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
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
