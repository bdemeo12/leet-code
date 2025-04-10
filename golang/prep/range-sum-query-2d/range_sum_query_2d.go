package main

import "fmt"

// Given a 2D matrix matrix, handle multiple queries of the following types:

// Update the value of a cell in matrix.

// Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

// Implement the NumMatrix class:

// NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.

// void update(int row, int col, int val) Updates the value of matrix[row][col] to be val.

// int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements
// of matrix inside the rectangle defined by its upper left corner
// (row1, col1) and lower right corner (row2, col2).

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	this.matrix[row][col] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.matrix[i][j]
		}
	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	matrix := make([][]int, 6)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 5)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
	obj := Constructor(matrix)

	fmt.Println(obj.SumRegion(0, 0, 5, 4))

}
