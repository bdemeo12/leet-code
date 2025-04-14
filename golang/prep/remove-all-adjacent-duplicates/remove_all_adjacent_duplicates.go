package main

import (
	"fmt"
)

// You are given a string s and an integer k, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them, causing the left and the right side of the deleted substring to concatenate together.

// We repeatedly make k duplicate removals on s until we no longer can.

// Return the final string after all such duplicate removals have been made. It is guaranteed that the answer is unique.

func removeDuplicates(s string, k int) string {
	char_stack := []string{}
	counter_stack := []int{}

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		// if the stack is not empty, and the top element is the same as curr
		if len(char_stack) > 0 && char_stack[len(char_stack)-1] == char {
			counter_stack[len(counter_stack)-1]++

			if counter_stack[len(counter_stack)-1] == k {
				// remove char
				char_stack = char_stack[:len(char_stack)-1]
				// remove counter
				counter_stack = counter_stack[:len(counter_stack)-1]
			}

		} else {
			// add char to stack
			char_stack = append(char_stack, char)
			counter_stack = append(counter_stack, 1)
		}
	}

	clean_string := ""
	for i := 0; i < len(char_stack); i++ {
		sub_string := ""
		for j := 0; j < counter_stack[i]; j++ {
			sub_string += char_stack[i]
		}

		clean_string += sub_string
	}

	return clean_string
}

func main() {
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3)) //aa
}
