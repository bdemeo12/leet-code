package main

import (
	"fmt"
	"sort"
)

// Two strings are considered close if you can attain one from the other using the following operations:

// Operation 1: Swap any two existing characters.
// For example, abcde -> aecdb
// Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
// For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
// You can use the operations on either string as many times as necessary.

// Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	map1 := stringToMap(word1)
	map2 := stringToMap(word2)

	if len(map1) != len(map2) {
		return false
	}

	return canSwap(map1, map2) || canTransform(map1, map2)
}

func canSwap(map1, map2 map[rune]int) bool {

	for key, val1 := range map1 {
		val2, ok := map2[key]

		if !ok || val1 != val2 {
			return false
		}
	}

	return true
}

func canTransform(map1, map2 map[rune]int) bool {

	var freq1 []int
	var chars1 []int
	for char, val := range map1 {
		freq1 = append(freq1, val)
		chars1 = append(chars1, int(char))
	}

	var freq2 []int
	var chars2 []int
	for char, val := range map2 {
		freq2 = append(freq2, val)
		chars2 = append(chars2, int(char))
	}

	chars1 = sorted(chars1)
	chars2 = sorted(chars2)

	for i := 0; i < len(chars1); i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	freq1 = sorted(freq1)
	freq2 = sorted(freq2)

	for i := 0; i < len(freq1); i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func sorted(x []int) []int {
	sort.Slice(x, func(a, b int) bool {
		return x[a] < x[b]
	})

	return x
}

func stringToMap(s string) map[rune]int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	return m
}

func main() {
	word1 := "uau"
	word2 := "ssx"
	fmt.Println(closeStrings(word1, word2))
}
