package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums and an integer k, return the k most frequent elements within the array.

// The test cases are generated such that the answer is always unique.

// You may return the output in any order.

func frequentElements(arr []int, k int) []int {

	elementCount := make(map[int]int)

	// add elements into map to keep count of frequency
	for _, element := range arr {
		if val, ok := elementCount[element]; ok {
			elementCount[element] = val + 1
			continue
		}

		elementCount[element] = 1
	}

	// sort
	kvPairs := sortMap(elementCount)

	// only want k elements
	var mostFrequent []int
	for i := 0; i < len(kvPairs[:k]); i++ {
		mostFrequent = append(mostFrequent, kvPairs[i].key)
	}

	return mostFrequent
}

func main() {
	fmt.Println(frequentElements([]int{1, 2, 2, 3, 3, 3}, 2))
}

type kv struct {
	key   int
	value int
}

func sortMap(m map[int]int) []kv {
	var pair []kv
	for k, v := range m {
		pair = append(pair, kv{k, v})
	}

	return sortArr(pair)
}

func sortArr(arr []kv) []kv {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].value > arr[j].value
	})

	return arr
}
