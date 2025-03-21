package main

import (
	"fmt"
)

// You are given a string s, which contains stars *.

// In one operation, you can:

// Choose a star in s.
// Remove the closest non-star character to its left, as well as remove the star itself.
// Return the string after all stars have been removed.

// Note:

// The input will be generated such that the operation is always possible.
// It can be shown that the resulting string will always be unique.

func removeStars(s string) string {

	stack := []rune{}

	for _, char := range s {
		if char == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}

func main() {
	s := "leet**cod*e" // Output: "lecoe"
	//s := "erase*****"
	fmt.Println(removeStars(s))
}
