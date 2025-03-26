package main

import (
	"fmt"
	"sort"
)

// You are given a 2D integer grid of size m x n and an integer x.
// In one operation, you can add x to or subtract x from any element in the grid.

// A uni-value grid is a grid where all the elements of it are equal.

// Return the minimum number of operations to make the grid uni-value.
// If it is not possible, return -1.

func minOperations(grid [][]int, x int) int {

	// flatten grid to 1d array

	flatGrid := []int{}
	remainder := grid[0][0] % x

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j]%x != remainder { // check that all number in the grid have the same remainder, if they dont its not possible
				return -1
			}

			flatGrid = append(flatGrid, grid[i][j])
		}
	}

	sort.Slice(flatGrid, func(a, b int) bool {
		return flatGrid[a] < flatGrid[b]
	})

	median := flatGrid[len(flatGrid)/2]
	operationCount := 0
	for i := 0; i < len(flatGrid); i++ {
		if flatGrid[i] > median { // subtract x
			for flatGrid[i] != median {
				flatGrid[i] -= x
				operationCount++
			}
		} else if flatGrid[i] < median { // add x
			for flatGrid[i] != median {
				flatGrid[i] += x
				operationCount++
			}
		} else if flatGrid[i] == median { // they are equal, move on
			continue
		}
	}

	return operationCount
}

func main() {
	grid := [][]int{{2, 4}, {6, 8}}
	x := 2
	fmt.Print(minOperations(grid, x))
}
