package main

import "fmt"

// Given an integer array nums,
// return true if you can partition the array into two subsets
// such that the sum of the elements in both subsets is equal or false otherwise.

// if arr sum is odd

func canPartition(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum%2 != 0 {
		return false // sum has to be even, no two subset can equal eachother if sum is odd. ex: 21 + 21, 11+ 11, 1+ 1
	}

	target := sum / 2

	canPartition := make([]bool, target+1)
	canPartition[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			canPartition[i] = canPartition[i] || canPartition[i-num]
		}
	}

	return canPartition[target]
}

func main() {
	nums := []int{1, 1, 5, 5}
	fmt.Println(canPartition(nums))

}
