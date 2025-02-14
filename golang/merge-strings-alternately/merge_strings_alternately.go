package main

import "strings"

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

func mergeAlternately(word1 string, word2 string) string {
	var mergedString string

	word1arr := strings.Split(word1, "")
	word2arr := strings.Split(word2, "")

	for i := 0; i < len(word1+word2); i++ {
		if i < len(word1arr) {
			mergedString = mergedString + word1arr[i]
		}

		if i < len(word2arr) {
			mergedString = mergedString + word2arr[i]
		}
	}

	return mergedString
}

func main() {
	// Example usage
	word1 := "abc"
	word2 := "defgh"
	result := mergeAlternately(word1, word2)
	println(result) // Output: "adbcefgh"
}
