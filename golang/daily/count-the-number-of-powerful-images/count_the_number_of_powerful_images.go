package main

import (
	"fmt"
	"strconv"
)

// You are given three integers start, finish, and limit.
//  You are also given a 0-indexed string s representing a positive integer.

// A positive integer x is called powerful if it ends with s (in other words, s is a suffix of x)
//  and each digit in x is at most limit.

// Return the total number of powerful integers in the range [start..finish].

// A string x is a suffix of a string y if and only if x is a substring of y
//  that starts from some index (including 0) in y and extends to the index y.length - 1.
//  For example, 25 is a suffix of 5125 whereas 512 is not.

// Digit DP is a technique used to solve problems that ask you to find the number of integers within a range that satisfies some property based on the digits of the integers.

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {

	// convert s string to array
	sArr := []int{}
	for _, char := range s {
		num, _ := strconv.Atoi(string(char))
		sArr = append(sArr, num)
	}

	// if a num in s is greater than the limit, then this will never work
	valid := true
	for _, n := range sArr {
		if n > limit {
			valid = false
		}
	}
	if !valid {
		return 0
	}

	powerfulInts := 0
	// a number cant be powerful if it is less than s
	sNum, _ := strconv.Atoi(s)
	for num := int64(sNum); num <= finish; num++ {

		// convert num from 123 to [1,2,3]
		digits := []int{}
		for num > 0 {
			digits = append([]int{int(num) % 10}, digits...) // prepend to digits slice
			num = num / 10
		}

		// check if suffix matches
		matches := true
		if len(digits) >= len(sArr) {
			for i := 0; i < len(sArr); i++ {
				if digits[len(digits)-len(sArr)+i] != sArr[i] {
					matches = false
					break
				}
			}
		}

		if matches {
			valid := true
			for _, d := range digits {
				if d > limit {
					valid = false
					break
				}
			}

			if valid {
				powerfulInts++
			}
		}
	}

	return int64(powerfulInts)
}

func main() {

	var start int64 = 1
	var finish int64 = 6000
	limit := 4
	s := "124"

	fmt.Println(numberOfPowerfulInt(start, finish, limit, s)) // output: 5
}
