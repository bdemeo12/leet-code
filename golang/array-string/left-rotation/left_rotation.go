package main

import "fmt"

// A left rotation operation on an array shifts each of the array's elements  unit to the left.
// For example, if 2 left rotations are performed on array 1,2,3,4,5 ,
//  then the array would become 3,4,5,1,2.
//  Note that the lowest index item moves to the highest index in a rotation. This is called a circular array.

// Given an array a of n integers and a number, d , perform d left rotations on the array.
// Return the updated array to be printed as a single line of space-separated integers.

// int a[n]: the array to rotate
// int d: the number of rotations

func rotLeft(a []int32, d int32) []int32 { // o(n) complexity
	n := len(a)
	d = d % int32(n) // handle if d > n
	return append(a[d:], a[:d]...)
}

// func rotLeft(a []int32, d int32) []int32 {

// 	var counter int32 = 0

// 	for counter < d {

// 		// remove from front
// 		tmp := a[0]

// 		a = append(a[:0], a[0+1:]...)

// 		// append it back
// 		a = append(a, tmp)

// 		counter++
// 	}
// 	return a
// }

func main() {

	input := []int32{1, 2, 3, 4, 5}
	fmt.Println(rotLeft(input, 2))
}
