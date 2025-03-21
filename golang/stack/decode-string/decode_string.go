package main

import (
	"fmt"
	"strconv"
)

// Given an encoded string, return its decoded string.

// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets
//  is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

// You may assume that the input string is always valid; there are no extra white spaces,
//  square brackets are well-formed, etc. Furthermore,
// you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k.
//  For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 105.

func decodeString(s string) string {
	stack := []rune{}
	var output string
	for _, char := range s {
		if char != ']' {
			stack = append(stack, char)
		} else {
			var curr string
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				curr = string(stack[len(stack)-1]) + curr
				stack = stack[:len(stack)-1] // pop
			}

			stack = stack[:len(stack)-1] // pop [

			var num string
			for len(stack) > 0 && isNum(stack[len(stack)-1]) {
				num = string(stack[len(stack)-1]) + num
				stack = stack[:len(stack)-1] // pop
			}
			times, _ := strconv.Atoi(num)

			repeated := ""
			for i := 0; i < times; i++ {
				repeated += curr
			}

			for _, ch := range repeated {
				stack = append(stack, ch)
			}
		}
	}

	return output
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func main() {
	//s := "3[a]2[bc]" // -> aaabcbc
	s := "3[a2[c]]" // -> accaccacc
	// s := "2[abc]3[cd]ef" // -> abcabccdcdcdef
	fmt.Println(decodeString(s))
}
