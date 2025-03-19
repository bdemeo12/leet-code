package main

import "fmt"

// There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes.
// // The biker starts his trip on point 0 with altitude equal 0.

// You are given an integer array gain of length n where gain[i] is the net gain in altitude between
// // points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.

func largestAltitude(gain []int) int {

	highest := 0
	count := 0

	for i := 0; i < len(gain); i++ {
		count = count + gain[i]

		if count > highest {
			highest = count
		}
	}

	return highest
}

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}
