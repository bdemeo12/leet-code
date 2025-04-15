package main

import "fmt"

// You are given two strings of the same length s and t.
//  In one step you can choose any character of t and replace it with another character.

// Return the minimum number of steps to make t an anagram of s.

// An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

func minSteps(s string, t string) int {

	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		tChar := t[i]
		tIndex := tChar - 'a' //db index in ascii
		count[tIndex]++

		sChar := s[i]
		sIndex := sChar - 'a'
		count[sIndex]--
	}

	sum := 0
	for _, c := range count {
		if c > 0 {
			sum += c
		}
	}

	return sum
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func main() {
	// s := "bab" //b2 a1
	// t := "aba" //a2 b1
	s := "leetcode"
	t := "practice" // 5
	fmt.Println(minSteps(s, t))
}
