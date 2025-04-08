package main

import "fmt"

// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
//  The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
//  The robot can only move either down or right at any point in time.

// Given the two integers m and n, return the number of possible unique paths that
// the robot can take to reach the bottom-right corner.

// The test cases are generated so that the answer will be less than or equal to 2 * 109.

func uniquePaths(m int, n int) int {

	// make grid
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	// first rows are 1, because it will always take one to get here
	for i := 0; i < m; i++ {
		grid[i][0] = 1
	}

	for j := 0; j < n; j++ {
		grid[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
