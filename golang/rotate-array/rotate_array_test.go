package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateArray(t *testing.T) {

	type TestRotateArrayParams struct {
		nums           []int
		k              int
		expectedResult []int
	}

	cases := []*TestRotateArrayParams{
		{
			nums:           []int{1, 2, 3, 4, 5, 6, 7},
			k:              3,
			expectedResult: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:           []int{-1, -100, 3, 99},
			k:              2,
			expectedResult: []int{3, 99, -1, -100},
		},
	}

	for _, c := range cases {
		result := rotateArray(c.nums, c.k)
		require.Equal(t, c.expectedResult, result)
	}
}
