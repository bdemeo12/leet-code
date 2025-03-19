package main

import "fmt"

// Given an array of integers nums and an integer target, return the indices i and j such that nums[i] + nums[j] == target and i != j.

// You may assume that every input has exactly one pair of indices i and j that satisfy the condition.

// Return the answer with the smaller index first.

func twoSum(arr []int, target int) []int {
	if len(arr) < 2 {
		return nil
	}

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if (arr[i] + arr[j]) == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{5, 5}, 10))
}
