package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func main() {
	type testCase struct {
		email string
		count int
	}

	runCases := []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
	}

	submitCases := append(runCases, []testCase{
		{"lila@bates.com", 31},
		{"sam@bates.com", 453},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.RWMutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		sc.mu.RLock()
		defer sc.mu.RUnlock()
		if output := sc.val(test.email); output != test.count {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
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
