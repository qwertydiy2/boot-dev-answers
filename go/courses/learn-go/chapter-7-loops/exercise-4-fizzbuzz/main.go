package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
