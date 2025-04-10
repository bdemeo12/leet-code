package main

import "fmt"

// You are given an m x n integer matrix grid and an array queries of size k.

// Find an array answer of size k such that for each integer queries[i]
//  you start in the top left cell of the matrix and repeat the following process:

// If queries[i] is strictly greater than the value of the current cell that you are in,
//  then you get one point if it is your first time visiting this cell, and you can move to any adjacent cell in all 4 directions: up, down, left, and right.

// Otherwise, you do not get any points, and you end this process.
// After the process, answer[i] is the maximum number of points you can get. Note that for each query you are allowed to visit the same cell multiple times.

// Return the resulting array answer.

// use dfs ?
func maxPoints(grid [][]int, queries []int) []int {
	var answer []int
	for _, query := range queries {

		// now we loop through the grid
		var points int
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {

				if query > grid[i][j] {
					points++ // increase point

					// check up
					// check down
					// check left
					// check right
				}

			}
		}
	}
	return answer
}

func main() {
	grid := [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}
	queries := []int{5, 6, 2}
	fmt.Println(maxPoints(grid, queries))
}
