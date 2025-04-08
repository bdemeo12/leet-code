package main

import "fmt"

// In this challenge, you need to calculate and print the sum of elements in an array,
//  considering that some integers may be very large.

// The first line of the input consists of an integer .
// The next line contains  space-separated integers contained in the array.

// Return the integer sum of the elements in the array.

// Contraints
// 1 <= n <= 10 -> only takes 10 elements
// 0 <= arr[i] <= 10^10 -> but elements can be very big

func bigSum(nums []int64) int64 {

	var sum int64
	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	nums := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(bigSum(nums))
}
