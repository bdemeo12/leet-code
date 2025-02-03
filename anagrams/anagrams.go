package main

import (
	"fmt"
	"strings"
)

// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

// An anagram is a string that contains the exact same characters as another string,
// but the order of the characters can be different.

func isAnagram(str1, str2 string) bool {

	if str1 == str2 {
		return true
	}

	if len(str1) != len(str2) {
		return false
	}

	str1Map := stringToMap(str1)
	str2Map := stringToMap(str2)

	// easy way to do it using reflect package, could also use maps.Equal()
	// return reflect.DeepEqual(str1Map, str2Map)

	for key, val := range str1Map {
		value, ok := str2Map[key]
		if ok && val == value {
			continue
		}

		return false
	}

	return true
}

func stringToMap(s string) map[string]int {
	strArr := strings.Split(s, "")
	strMap := make(map[string]int)
	for _, char := range strArr {
		val := strMap[char]
		strMap[char] = val + 1
	}

	return strMap
}

func main() {
	_ = isAnagram("bat", "hel")
	fmt.Println("done")
}
