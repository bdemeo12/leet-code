package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	groupedAnagrams := make(map[string][]string)

	for _, str := range strs {
		key := sortedString(str)
		groupedAnagrams[key] = append(groupedAnagrams[key], str)
	}

	var result [][]string
	for _, val := range groupedAnagrams {
		result = append(result, val)
	}

	return result
}

func sortedString(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool { // convert each char into a int, and sort based on ints
		return r[i] < r[j]
	})

	return string(r)
}

func main() {
	fmt.Println(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
}
