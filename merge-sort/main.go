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

func mergeSort(nums1 []int, m int, nums2 []int, n int) []int {
	// If there is just 1 item in each array
	if len(nums1) <= 1 && len(nums2) <= 1 {
		if nums1[0] < nums2[0] { // sort in non decreasing order aka increasing order
			nums1 = append(nums1, nums2[0])
		} else {
			nums1 = append([]int{nums2[0]}, nums1...) // is this the best way ?
		}

		// confirm the array len is correct before we finish
		if len(nums1) != (m + n) {
			nums1 = nums1[:m+n]
		}

		return nums1
	}

	nums1 = divideAndConquer(nums1)
	nums2 = divideAndConquer(nums2)

	nums1 = append(nums1, nums2...)
	nums1 = divideAndConquer(nums1)
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

	fmt.Println("final result:")

	return result
}

func main() {
	nums1 := []int{4, 3, 5}
	m := 3
	fmt.Println("Nums1:", nums1, "m:", m)

	nums2 := []int{6, 0}
	n := 1
	fmt.Println("Nums2:", nums2, "n:", n)

	fmt.Println(mergeSort(nums1, m, nums2, n))
}
