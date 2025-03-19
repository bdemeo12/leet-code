package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicatesii(t *testing.T) {

	type removeDuplicatesParamsii struct {
		nums []int

		expectedResult int
	}

	cases := []*removeDuplicatesParamsii{
		{
			nums:           []int{1, 1, 1, 2, 2, 3},
			expectedResult: 5,
		},
		{
			nums:           []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedResult: 7,
		},
	}

	for _, c := range cases {
		result := removeDuplicatesii(c.nums)
		require.Equal(t, c.expectedResult, result)
	}
}
