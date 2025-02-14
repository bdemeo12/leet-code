package main

import "fmt"

// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

func maxArea(height []int) int {

	i, j := 0, len(height)-1
	var maxArea int
	for i != j {
		// area is calculated with l x w
		width := j - i
		area := 0

		if height[i] <= height[j] {
			area = height[i] * width
			i++
		} else if height[i] > height[j] {
			area = height[j] * width
			j--
		}

		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5}))
}
