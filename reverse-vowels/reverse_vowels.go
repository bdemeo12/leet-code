package main

import "strings"

// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower
// and upper cases, more than once.

func reverseVowels(s string) string {

	vowels := "aeiouAEIOU"

	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		if !strings.Contains(vowels, string(s[i])) {
			i++
			continue
		}

		if !strings.Contains(vowels, string(s[j])) {
			j--
			continue
		}
		runes[i], runes[j] = runes[j], runes[i]

		i++
		j--
	}
	return string(runes)
}
