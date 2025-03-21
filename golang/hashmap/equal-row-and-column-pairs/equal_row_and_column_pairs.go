package main

import (
	"fmt"
	"strconv"
)

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

func equalPairs(grid [][]int) int {
	// Its not possible to have equal pairs, if the row length and col length are different
	if len(grid) != len(grid[0]) {
		return 0
	}

	rowCount := make(map[string]int)
	colCount := make(map[string]int)

	for i := 0; i < len(grid); i++ { // rows

		// // add the row to the map
		row := arrToString(grid[i])
		rowCount[row]++

		col := ""
		for j := 0; j < len(grid[i]); j++ { // cols
			col = col + strconv.Itoa(grid[j][i]) + ","

			if j == len(grid[i])-1 { // we have reached the end of the col, we can add to the map
				colCount[col]++
			}
		}
	}

	// now lets iterate though the map, if the count is greater than 1, count pair
	pairs := 0
	for key, rowCount := range rowCount {
		if colCount, ok := colCount[key]; ok {
			pairs += rowCount * colCount
		}
	}

	return pairs
}

func arrToString(arr []int) string {
	var s string
	for _, val := range arr {
		s = s + strconv.Itoa(val) + ","
	}

	return s
}

func main() {

	//grid := [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}
	//grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	grid := [][]int{{13, 13}, {13, 13}}
	//grid := [][]int{{11, 1}, {1, 11}}
	fmt.Println(equalPairs(grid))

}
