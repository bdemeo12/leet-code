package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {

	type removeDuplicatesParams struct {
		nums []int

		expectedResult int
	}

	cases := []*removeDuplicatesParams{
		{
			nums:           []int{1, 1, 2},
			expectedResult: 2,
		},
		{
			nums:           []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedResult: 5,
		},
	}

	for _, c := range cases {
		result := removeDuplicates(c.nums)
		require.Equal(t, c.expectedResult, result)
	}
}
