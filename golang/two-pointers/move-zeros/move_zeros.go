package main

// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

func moveZeroes(nums []int) {
	lastNonZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Swap current element with lastNonZero position
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}

func main() {

	// Input: nums = [0,1,0,3,12]
	// Output: [1,3,12,0,0]
	//nums := []int{0, 0, 1}
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)
}
