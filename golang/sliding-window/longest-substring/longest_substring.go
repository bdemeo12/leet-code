package main

// Given a string s and an integer k, return the length of the longest substring of s such
// that the frequency of each character in this substring is greater than or equal to k.

// if no such substring exists, return 0.

// sliding window

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	return 0
}

func main() {
	s := "aaabb"
	k := 3
	result := longestSubstring(s, k)
	println(result) // Output: 3
}
