package main

import "fmt"

// You are given two positive integers low and high.

// An integer x consisting of 2 * n digits is symmetric if the sum of the first n digits of x
// is equal to the sum of the last n digits of x. Numbers with an odd number of digits are never symmetric.

// Return the number of symmetric integers in the range [low, high].

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for num := low; num <= high; num++ {
		digits := splitDigits(num)

		if len(digits)%2 != 0 {
			continue
		}

		n := len(digits) / 2

		sum1 := 0
		for i := 0; i < n; i++ {
			sum1 += digits[i]
		}

		sum2 := 0
		for j := n; j < len(digits); j++ {
			sum2 += digits[j]
		}

		if sum1 == sum2 {
			count++
		}
	}

	return count
}

func splitDigits(n int) []int {
	digits := []int{}
	for _, ch := range fmt.Sprintf("%d", n) {
		digits = append(digits, int(ch-'0'))
	}
	return digits
}

func main() {
	low := 1
	high := 100
	fmt.Println(countSymmetricIntegers(low, high))
}
