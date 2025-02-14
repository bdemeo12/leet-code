package main

import "fmt"

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

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(rotateArray(nums, k))
}
