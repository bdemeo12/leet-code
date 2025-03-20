package main

import "fmt"

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

// Note that the integers in the lists may be returned in any order.

func findDifference(nums1 []int, nums2 []int) [][]int {

	// convert arrays to maps of nums to count

	nums1map := make(map[int]bool)

	for _, num := range nums1 {
		nums1map[num] = true
	}

	nums2map := make(map[int]bool)

	for _, num := range nums2 {

		nums2map[num] = true
	}

	// now find unique values
	var arr1 []int
	for key, _ := range nums1map {
		if _, ok := nums2map[key]; !ok {
			arr1 = append(arr1, key)
		}

	}

	// do the same for nums 2

	var arr2 []int
	for key, _ := range nums2map {
		if _, ok := nums1map[key]; !ok {
			arr2 = append(arr2, key)
		}
	}

	return [][]int{arr1[:], arr2[:]}
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}

	fmt.Println(findDifference(nums1, nums2))
}
