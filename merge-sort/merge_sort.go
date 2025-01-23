package main

import (
	"fmt"
)

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be
// merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// ************************
// Question: why can we not combine nums1 and nums 2 off the bat ? is there a change in complexity if we process them together, vs process them seperatly and then remerge and then process ? Seems like the same thing?
// Question: is it faster/better to remove all 0s from the arrays before we start ?
// ************************

// func mergeSort(nums1 []int, m int, nums2 []int, n int) []int {
// 	// empty array check
// 	if len(nums1) == 0 || len(nums2) == 0 {
// 		nums1 = append(nums1, nums2...)
// 		return nums1
// 	}

// 	nums1 = divideAndConquer(nums1)
// 	nums2 = divideAndConquer(nums2)

// 	nums1 = append(nums1, nums2...)
// 	nums1 = divideAndConquer(nums1)
// 	// trim
// 	arrayLen := m + n
// 	nums1 = nums1[len(nums1)-arrayLen:]

// 	return nums1

// }

func mergeSort(nums1 []int, m int, nums2 []int, n int) []int {
	// empty array check
	if len(nums1) == 0 || len(nums2) == 0 {
		nums1 = append(nums1, nums2...)
		return nums1
	}

	// lets deal with one array
	nums1 = divideAndConquer(append(nums1, nums2...))

	// trim
	arrayLen := m + n
	nums1 = nums1[len(nums1)-arrayLen:]

	return nums1

}

func divideAndConquer(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	half := len(arr) / 2
	left := divideAndConquer(arr[:half])
	right := divideAndConquer(arr[half:])

	return merge(left, right)
}

func merge(left, right []int) []int {

	i := 0
	j := 0
	result := []int{}

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	var nums1 []int
	m := 0
	fmt.Println("Nums1:", nums1, "m:", m)

	nums2 := []int{1}
	//var nums2 []int
	n := 1
	fmt.Println("Nums2:", nums2, "n:", n)

	fmt.Println(mergeSort(nums1, m, nums2, n))
}
