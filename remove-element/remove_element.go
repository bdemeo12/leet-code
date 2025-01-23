package main

import "fmt"

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	k := len(nums)

	for i := 0; i < k; i++ {
		if nums[i] == val {
			tmp := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, tmp)
		}
	}

	fmt.Println("output array:", nums)
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	_ = removeElement(nums, val)
}
