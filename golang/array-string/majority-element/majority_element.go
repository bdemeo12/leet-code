package main

import "fmt"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

func majorityElement(nums []int) int {

	if len(nums) == 0 {
		return -1 // or any other value that indicates no majority element found
	}

	for _, n := range nums {
		count := 0
		for _, m := range nums {
			if n == m {
				count++
			}
		}

		if count > len(nums)/2 {
			return n
		}
	}

	return -1 // or any other value that indicates no majority element found

}

func main() {

	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
