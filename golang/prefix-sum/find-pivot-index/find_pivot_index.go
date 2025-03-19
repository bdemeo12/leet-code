package main

import "fmt"

// Given an array of integers nums, calculate the pivot index of this array.

// The pivot index is the index where the sum of all the numbers strictly to the
//  left of the index is equal to the sum of all the numbers strictly to the index's right.

// If the index is on the left edge of the array, then the left sum is 0
// because there are no elements to the left. This also applies to the right edge of the array.

// Return the leftmost pivot index. If no such index exists, return -1.

func pivotIndex(nums []int) int {

	for i := 0; i < len(nums); i++ {

		left := 0
		for j := 0; j < i; j++ {
			left = left + nums[j]
		}

		right := 0
		for k := len(nums) - 1; k > i; k-- {
			right = right + nums[k]
		}

		if left == right {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(pivotIndex(nums))
}
