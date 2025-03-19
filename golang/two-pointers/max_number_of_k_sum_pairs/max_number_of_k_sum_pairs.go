package main

import (
	"fmt"
	"sort"
)

// You are given an integer array nums and an integer k.

// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

// Return the maximum number of operations you can perform on the array.

func maxOperations(nums []int, k int) int {

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	max := 0

	left, right := 0, len(nums)-1

	for left < right {

		sum := nums[left] + nums[right]

		if sum == k {
			max++
			left++
			right--
		} else if sum > k {
			right--
		} else if sum < k {
			left++
		}
	}

	return max
}

func main() {
	fmt.Println(maxOperations([]int{2, 2, 2, 3, 1, 1, 4, 1}, 4))
}

// func maxOperations(nums []int, k int) int {

// 	x := make(map[int]int)
// 	var operations int
// 	for i := 0; i < len(nums); i++ {
// 		if val, ok := x[nums[i]]; !ok {
// 			x[nums[i]] = 1
// 		} else {
// 			x[nums[i]] = val + 1
// 		}
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		pairVal := k - nums[i]

// 		if val, ok := x[pairVal]; ok {
// 			val = val - 1
// 			x[pairVal] = val
// 			operations++

// 			if val == 0 {
// 				delete(x, nums[i])
// 				delete(x, pairVal)
// 			}
// 		}
// 	}
// 	return operations
// }
