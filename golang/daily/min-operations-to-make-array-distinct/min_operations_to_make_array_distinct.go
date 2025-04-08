package main

import "fmt"

// You are given an integer array nums.
//  You need to ensure that the elements in the array are distinct.
//  To achieve this, you can perform the following operation any number of times:

// Remove 3 elements from the beginning of the array.
// If the array has fewer than 3 elements, remove all remaining elements.

// Note that an empty array is considered to have distinct elements.
// Return the minimum number of operations needed to make the elements in the array distinct.

// Notes
// i / 3 + 1

func minimumOperations(nums []int) int {

	seen := make(map[int]bool)

	for i := len(nums) - 1; i >= 0; i-- { // reverse iteration
		if ok := seen[nums[i]]; ok {
			return i/3 + 1 // i / 3 determins how many group of element can be removed up to index i, + 1 accounts for the operation that is doing the removal now
		}

		seen[nums[i]] = true
	}

	return 0
}

// func minimumOperations(nums []int) int {

// 	// check if elements are distinct
// 	distinct := make(map[int]int) // map of number to count
// 	for _, num := range nums {
// 		distinct[num]++
// 	}

// 	count := 0
// 	for _, val := range distinct {
// 		if val > 1 {
// 			if len(nums) <= 3 {
// 				return 1 // removing 3 or fewer takes 1 operation
// 			}

// 			// remove first three elements
// 			nums = nums[3:]
// 			count = 1 + minimumOperations(nums)
// 			return count
// 		}
// 	}

// 	return count
// }

func main() {
	nums := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
	//nums := []int{15, 3, 5, 5}
	fmt.Println(minimumOperations(nums))
}
