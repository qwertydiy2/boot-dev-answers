package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func fizzbuzz() {
	// Implementation goes here
	// Ensure all output is lowercase
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("fizzbuzz\n")
		} else if i%5 == 0 {
			fmt.Printf("buzz\n")
		} else if i%3 == 0 {
			fmt.Printf("fizz\n")
		}
	}
}

// Helper used by both main and TestFizzBuzzOutput
func checkFizzBuzzOutput() error {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call fizzbuzz
	fizzbuzz()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Only multiples of 3 or 5 should produce output, so lines may be fewer than 100
	// Let's collect the expected outputs for 1..100
	expected := []string{}
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			expected = append(expected, "fizzbuzz")
		} else if i%3 == 0 {
			expected = append(expected, "fizz")
		} else if i%5 == 0 {
			expected = append(expected, "buzz")
		}
		// else: no output
	}

	// Check line count matches expected output count
	if len(lines) != len(expected) {
		return fmt.Errorf("Expected %d lines, got %d", len(expected), len(lines))
	}

	// Compare outputs
	for idx, exp := range expected {
		if strings.ToLower(lines[idx]) != exp {
			return fmt.Errorf("line %d: expected %q, got %q", idx+1, exp, strings.ToLower(lines[idx]))
		}
	}

	// Count occurrences for completeness
	fizzCount, buzzCount, fizzbuzzCount := 0, 0, 0
	for _, line := range lines {
		switch strings.ToLower(line) {
		case "fizzbuzz":
			fizzbuzzCount++
		case "fizz":
			fizzCount++
		case "buzz":
			buzzCount++
		}
	}
	if fizzCount != 27 {
		return fmt.Errorf("Expected 27 fizz, got %d", fizzCount)
	}
	if buzzCount != 14 {
		return fmt.Errorf("Expected 14 buzz, got %d", buzzCount)
	}
	if fizzbuzzCount != 6 {
		return fmt.Errorf("Expected 6 fizzbuzz, got %d", fizzbuzzCount)
	}
	return nil
}

// don't touch below this line

func main() {
	fizzbuzz()
	if err := checkFizzBuzzOutput(); err != nil {
		fmt.Println("Test failed:", err)
	} else {
		fmt.Println("All tests passed!")
	}
}
