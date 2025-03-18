package main

import "fmt"

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below the provided parameter value number.

func multiplesOf3Or5(x int) int {

	sum := 0
	for i := 0; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Printf("%d is a multiple \n", i)
			sum += i
		}
	}

	return sum
}

func main() {
	fmt.Println(multiplesOf3Or5(1000))
}
