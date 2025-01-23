package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {

	type mergeSortParams struct {
		nums []int
		val  int

		expectedResult int
	}

	cases := []*mergeSortParams{
		{
			nums:           []int{3, 2, 2, 3},
			val:            3,
			expectedResult: 2,
		},
		{
			nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:            2,
			expectedResult: 5,
		},
	}

	for _, c := range cases {
		result := removeElement(c.nums, c.val)
		require.Equal(t, c.expectedResult, result)
	}
}
