package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 || i == 3 {
			fmt.Printf("%d\n", i)
		}
		if i%6 == 1 || i%6 == 5 {
			primeNumber := true
			for j := 2; j <= int(math.Ceil(math.Sqrt(float64(i)))); j++ {
				if i%j == 0 {
					primeNumber = false
					break
				}
			}
			if primeNumber {
				fmt.Printf("%d\n", i)
			}
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
