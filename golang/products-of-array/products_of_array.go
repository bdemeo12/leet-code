package main

import "fmt"

// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

func productsOfArray(arr []int) []int {
	var productArr []int

	for i := 0; i < len(arr); i++ {
		var tmpArr = make([]int, len(arr))
		copy(tmpArr, arr)                            // need to make a copy of arr so we dont modify it
		tmpArr = append(tmpArr[:i], tmpArr[i+1:]...) // create a tmp array with element at i removed
		val := 1
		for _, element := range tmpArr {
			val = val * element
		}

		productArr = append(productArr, val)
	}

	return productArr
}

func main() {
	fmt.Println(productsOfArray([]int{1, 2, 4, 6}))
}
