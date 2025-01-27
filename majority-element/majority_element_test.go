package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMajorityElement(t *testing.T) {

	type majorityElementsParams struct {
		nums []int

		expectedResult int
	}

	cases := []*majorityElementsParams{
		{
			nums:           []int{3, 2, 3},
			expectedResult: 3,
		},
		{
			nums:           []int{2, 2, 1, 1, 1, 2, 2},
			expectedResult: 2,
		},
	}

	for _, c := range cases {
		result := majorityElement(c.nums)
		require.Equal(t, c.expectedResult, result)
	}
}
