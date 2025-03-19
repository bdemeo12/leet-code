package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	type twoSumParams struct {
		arr            []int
		target         int
		expectedResult []int
	}

	cases := []twoSumParams{
		{
			arr:            []int{3, 4, 5, 6},
			target:         7,
			expectedResult: []int{0, 1},
		},
		{
			arr:            []int{4, 5, 6},
			target:         10,
			expectedResult: []int{0, 2},
		},
		{
			arr:            []int{5, 5},
			target:         10,
			expectedResult: []int{0, 1},
		},
		{
			arr:            []int{1},
			target:         100,
			expectedResult: nil,
		},
		{
			arr:            []int{1, 2, 3},
			target:         500,
			expectedResult: nil,
		},
	}

	for _, c := range cases {
		result := twoSum(c.arr, c.target)
		assert.Equal(t, c.expectedResult, result)
	}
}
