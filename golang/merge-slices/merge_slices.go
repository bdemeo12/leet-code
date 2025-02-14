package main

import (
	"fmt"
	"sort"
)

func mergeSlices(a, b []int) []int {

	a = sortSlice(a)
	b = sortSlice(b)

	i := 0
	var merged []int
	for i < (len(a) + len(b)) {
		if i < len(a) {
			merged = append(merged, a[i])
		}

		if i < len(b) {
			merged = append(merged, b[i])
		}

		i++
	}

	return merged
}

func sortSlice(s []int) []int {
	sort.Slice(s, func(a, b int) bool {
		return s[a] < s[b]
	})

	return s
}

func main() {
	b := []int{8, 6, 4, 2}
	a := []int{9, 7, 5, 3, 1}

	fmt.Println(mergeSlices(a, b))
}
