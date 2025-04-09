package main

import (
	"fmt"
	"sort"
)

// The power of an integer x is defined as the number of steps needed to transform x into 1 using the following steps:

// if x is even then x = x / 2
// if x is odd then x = 3 * x + 1
// For example, the power of x = 3 is 7 because 3 needs 7 steps to become 1 (3 --> 10 --> 5 --> 16 --> 8 --> 4 --> 2 --> 1).

// Given three integers lo, hi and k.
// The task is to sort all integers in the interval [lo, hi]
//  by the power value in ascending order, if two or more integers
// have the same power value sort them by ascending order.

// Return the kth integer in the range [lo, hi] sorted by the power value.

// Notice that for any integer x (lo <= x <= hi) it is guaranteed that x will transform into 1
//  using these steps and that the power of x is will fit in a 32-bit signed integer.

func getKth(lo int, hi int, k int) int {

	nums := [][]int{}
	for i := lo; i <= hi; i++ {
		nums = append(nums, []int{i, 0})
	}

	// find operations - save in 2d array of num/operations
	for i, num := range nums {

		count := 0
		tmp := num[0]
		for num[0] != 1 {
			count++
			if num[0]%2 == 0 {
				num[0] = num[0] / 2
			} else {
				num[0] = (3 * num[0]) + 1
			}
		}

		nums[i] = []int{tmp, count}
	}

	// sort by operation count
	sort.Slice(nums, func(a, b int) bool {
		if nums[a][1] == nums[b][1] {
			return nums[a][0] < nums[b][0]
		}
		return nums[a][1] < nums[b][1]
	})

	return nums[k-1][0]
}

func main() {
	lo := 1
	hi := 1000
	k := 777
	fmt.Println(getKth(lo, hi, k)) // 13
}
