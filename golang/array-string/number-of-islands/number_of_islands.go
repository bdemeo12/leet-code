package main

import "fmt"

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.

func numOfIslands(grid [][]byte) int {

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}

		}
	}

	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)

	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

func main() {
	grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Print(numOfIslands(grid))
}
