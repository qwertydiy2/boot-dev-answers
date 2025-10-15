package main

import (
	"fmt"
	"slices"
)

func concurrentFib(n int) []int {
	// ?
	results := make([]int, 0)
	fibCh := make(chan int)
	go fibonacci(n, fibCh)
	for result := range fibCh {
		results = append(results, result)
	}
	return results
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	type testCase struct {
		n        int
		expected []int
	}

	runCases := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}

	submitCases := append(runCases, []testCase{
		{0, []int{}},
		{1, []int{0}},
		{7, []int{0, 1, 1, 2, 3, 5, 8}},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
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
