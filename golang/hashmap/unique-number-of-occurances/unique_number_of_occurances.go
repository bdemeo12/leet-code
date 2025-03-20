package main

import "fmt"

// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

func uniqueOccurrences(arr []int) bool {

	valueCount := make(map[int]int)

	for _, a := range arr {
		if count, ok := valueCount[a]; !ok {
			valueCount[a] = 1
		} else {
			valueCount[a] = count + 1
		}
	}

	uniqueCount := make(map[int]int)
	for _, val := range valueCount {
		if count, ok := uniqueCount[val]; !ok {
			uniqueCount[val] = 1
		} else {
			uniqueCount[val] = count + 1
		}
	}

	for _, val := range uniqueCount {
		if val > 1 {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))
}
