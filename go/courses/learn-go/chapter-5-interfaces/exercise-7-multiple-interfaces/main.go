package main

import (
    "fmt"
)

type email struct {
    isSubscribed bool
    body         string
}

func (e email) cost() int {
    if e.isSubscribed {
        return len(e.body) * 2
    }
    return len(e.body) * 5
}

func (e email) format() string {
    var subscribedString string
    if e.isSubscribed {
        subscribedString = "Subscribed"
    } else {
        subscribedString = "Not Subscribed"
    }
    return fmt.Sprintf("'%s' | %s", e.body, subscribedString)
}

type expense interface {
    cost() int
}

type formatter interface {
    format() string
}

var withSubmit = true

func main() {
    type testCase struct {
        body           string
        isSubscribed   bool
        expectedCost   int
        expectedFormat string
    }

    runCases := []testCase{
        {"hello there", true, 22, "'hello there' | Subscribed"},
        {"general kenobi", false, 70, "'general kenobi' | Not Subscribed"},
    }

    submitCases := append(runCases, []testCase{
        {"i hate sand", true, 22, "'i hate sand' | Subscribed"},
        {"it's coarse and rough and irritating", false, 180, "'it's coarse and rough and irritating' | Not Subscribed"},
        {"and it gets everywhere", true, 44, "'and it gets everywhere' | Subscribed"},
    }...)

    testCases := runCases
    if withSubmit {
        testCases = submitCases
    }

    skipped := len(submitCases) - len(testCases)

    passCount := 0
    failCount := 0

    for _, test := range testCases {
        e := email{
            body:         test.body,
            isSubscribed: test.isSubscribed,
        }
        cost := e.cost()
        format := e.format()
        if format != test.expectedFormat || cost != test.expectedCost {
            failCount++
            fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
        } else {
            passCount++
            fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.body, test.isSubscribed, test.expectedCost, test.expectedFormat, cost, format)
        }
    }

    fmt.Println("---------------------------------")
    if skipped > 0 {
        fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
    } else {
        fmt.Printf("%d passed, %d failed\n", passCount, failCount)
    }
}