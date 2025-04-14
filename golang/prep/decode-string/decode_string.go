package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// Given an encoded string, return its decoded string.

// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets
//  is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

// You may assume that the input string is always valid; there are no extra white spaces,
//  square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 105.

// backtracking ?

func decodeString(s string) string {

	stack := []rune{}
	for _, char := range s {
		if char != ']' {
			// push to stack
			stack = append(stack, char)
			continue
		}

		decodedString := []rune{}
		last := stack[len(stack)-1]
		for last != '[' { // while the last item in stack is not[
			// add last to decoded string
			decodedString = append([]rune{last}, decodedString...)
			// remove last from stack
			stack = stack[:len(stack)-1]
			// update last
			last = stack[len(stack)-1]
		}

		// Pop the '['
		stack = stack[:len(stack)-1]

		// get times
		k := []rune{}
		for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {

			// push to k
			k = append([]rune{stack[len(stack)-1]}, k...)

			// pop digit
			stack = stack[:len(stack)-1]
		}

		times, _ := strconv.Atoi(string(k))

		repeated := []rune{}
		for range times {
			repeated = append(repeated, decodedString...)
		}

		stack = append(stack, repeated...)
	}

	return string(stack)
}

func main() {
	//s := "3[a]2[bc]"
	// s := "3[a2[c]]" //accaccacc
	//s := "2[abc]3[cd]ef" // abcabccdcdcdef
	s := "100[leetcode]"
	fmt.Println(decodeString(s))
}
