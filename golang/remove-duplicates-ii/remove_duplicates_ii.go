package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique
// element appears at most twice. The relative order of the elements should be kept the same.

// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in
// the first part of the array nums. More formally, if there are k elements after removing the duplicates,
// then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

// Return k after placing the final result in the first k slots of nums.

// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

func removeDuplicatesii(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// k is the position to place the next unique element
	k := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		// Compare the current element with the previous one
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[k] = nums[i] // Move the unique element to the front
			k++
		}
	}

	fmt.Println(nums[:k])
	return k
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicatesii(nums))
}
