package main

import "fmt"

// You are given an integer array nums consisting of n elements, and an integer k.

// Find a contiguous subarray whose length is equal to k that has the maximum average value and
// return this value. Any answer with a calculation error less than 10-5 will be accepted.

func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum

	for i := k; i < len(nums); i++ {
		// add new element
		sum += nums[i] - nums[i-k]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4

	fmt.Println("Average: ", findMaxAverage(nums, k))
}
