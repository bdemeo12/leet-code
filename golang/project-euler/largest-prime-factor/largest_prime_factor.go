package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the given number?

func largestPrimeFactor(x int) int {
	if x == 1 {
		return 2
	}

	if x%2 == 0 {
		x = x / 2
	}

	largestPrimeFactor := 2
	for i := 3; i <= x; i += 2 {
		if x%i == 0 {
			largestPrimeFactor = i
			x = x / i
		}
	}

	return largestPrimeFactor
}

func main() {
	fmt.Println(largestPrimeFactor(13195))
}
