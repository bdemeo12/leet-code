package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were
// present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.

func removeDuplicates(nums []int) int {
	uniqueNums := make(map[int]int)

	for i, n := range nums {
		// If we dont have it in the map, set it
		if _, ok := uniqueNums[n]; !ok {
			uniqueNums[n] = n
			continue
		}

		// else, we have it already, lets remove it from nums
		nums = append(nums[:i], nums[i+1:]...)
	}

	fmt.Println("nums:", nums)
	return len(nums)
}

func main() {
	nums := []int{3, 1, 1, 4, 2, 4}
	_ = removeDuplicates(nums)
}
