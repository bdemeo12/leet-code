package main

import "fmt"

// You are given a binary array nums.

// You can do the following operation on the array any number of times (possibly zero):

// Choose any 3 consecutive elements from the array and flip all of them.
// Flipping an element means changing its value from 0 to 1, and from 1 to 0.

// Return the minimum number of operations required to make all elements in nums
//  equal to 1. If it is impossible, return -1.

// sliding window
func minOperations(nums []int) int {

	count := 0

	for i := 2; i < len(nums); i++ { // start at the third position and look back
		if nums[i-2] == 0 {
			//flip

			for j := i - 2; j < i+1; j++ {
				if nums[j] == 0 {
					nums[j] = 1
				} else if nums[j] == 1 {
					nums[j] = 0
				}
			}
			count++
		}

	}

	// check if final arry contains a 0
	for _, num := range nums {
		if num == 0 {
			return -1
		}
	}

	return count
}

func main() {
	nums := []int{0, 1, 1, 1, 0, 0}
	// 1 0 0 1 0 0
	// 1 1 1 0 0 0
	// 1 1 1 0 0 0
	fmt.Println(minOperations(nums))
}
