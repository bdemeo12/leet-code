package main

// Given an integer array nums, return true if there exists a triple of indices (i, j, k)
// such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	smallest, smaller := 0, 0
	first, second := true, true

	for i := 0; i < len(nums); i++ {

		if first || nums[i] <= smallest {
			smallest = nums[i]
			first = false
		} else if second || nums[i] <= smaller {
			smaller = nums[i]
			second = false
		} else {
			return true
		}

	}
	return false
}
