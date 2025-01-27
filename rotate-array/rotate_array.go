package main

func rotateArray(nums []int, k int) []int {

	if len(nums) == 0 || k == 0 {
		return nums
	}

	var tempArr []int
	tempArr = append(tempArr, nums[len(nums)-k:]...)

	// now, prepend
	nums = append(tempArr, nums[:len(nums)-k]...)

	return nums
}
