package main

// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

func moveZeros(nums []int) []int {

	i, count := 0, 0
	for i < len(nums) {
		if count == len(nums) {
			break
		}

		if nums[i] == 0 {
			nums = append(nums, nums[i])

			// remove from the position in place
			nums = append(nums[:i], nums[i+1:]...)

			count++
			continue
		}

		count++
		i++
	}

	return nums
}
