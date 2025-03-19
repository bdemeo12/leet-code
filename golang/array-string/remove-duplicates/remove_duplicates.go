package main

import "fmt"

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were
// present in nums initially. The remaining elements of nums are not important as well as the size of nums.
// Return k.

// using 1 array
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// k is the position to place the next unique element
	k := 1

	for i := 1; i < len(nums); i++ {
		// Compare the current element with the previous one
		if nums[i] != nums[i-1] {
			nums[k] = nums[i] // Move the unique element to the front
			k++
		}
	}

	fmt.Println(nums[:k])
	return k

}

// using 2 arrays
// func removeDuplicates(nums []int) int {
// 	uniqueNums := make(map[int]int)

// 	for _, n := range nums {
// 		uniqueNums[n] = n
// 	}

// 	newNums := []int{}
// 	for k, _ := range uniqueNums {
// 		newNums = append(newNums, k)
// 	}

// 	fmt.Println("nums:", newNums)
// 	return len(newNums)
// }

func main() {
	nums := []int{1, 0, 523, 5, 0, 2, 2}
	_ = removeDuplicates(nums)
}
