package main

import "strconv"

// Given an array of characters chars, compress it using the following algorithm:

// Begin with an empty string s. For each group of consecutive repeating characters in chars:

// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately,
// but instead, be stored in the input character array chars. Note that group lengths that are 10 or
// longer will be split into multiple characters in chars.

// After you are done modifying the input array, return the new length of the array.

// You must write an algorithm that uses only constant extra space.

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	s := ""
	counter := 1

	for i := 1; i < len(chars); i++ {

		if chars[i] == chars[i-1] {
			counter++

			if i != len(chars)-1 {
				continue
			}
		}

		// if they are not equal we can append to s
		s = s + string(chars[i-1])
		if counter > 1 {
			s = s + strconv.Itoa(counter)
		}
		// reset counter
		counter = 1
	}
	return len(s)
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}

	result := compress(chars)
	println(result)

	// print the compressed array
	for i := 0; i < result; i++ {
		print(string(chars[i]))
	}
}
