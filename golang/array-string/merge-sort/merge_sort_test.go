package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {

	type mergeSortParams struct {
		nums1 []int
		nums2 []int
		m     int
		n     int

		expectedResult []int
	}

	cases := []*mergeSortParams{
		{
			nums1:          []int{1, 2, 3, 0, 0, 0},
			nums2:          []int{2, 5, 6},
			m:              3,
			n:              3,
			expectedResult: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:          []int{1},
			nums2:          []int{},
			m:              1,
			n:              0,
			expectedResult: []int{1},
		},
		{
			nums1:          []int{0},
			nums2:          []int{1},
			m:              0,
			n:              1,
			expectedResult: []int{1},
		},
	}

	for _, c := range cases {
		result := mergeSort(c.nums1, c.m, c.nums2, c.n)
		require.Equal(t, c.expectedResult, result)
	}
}
