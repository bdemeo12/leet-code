package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequentElements(t *testing.T) {

	type frequentElementsParams struct {
		arr            []int
		k              int
		expectedResult []int
	}

	cases := []frequentElementsParams{
		{
			arr:            []int{1, 2, 2, 3, 3, 3},
			k:              2,
			expectedResult: []int{2, 3},
		},
		{
			arr:            []int{7, 7},
			k:              1,
			expectedResult: []int{7},
		},
	}

	for _, c := range cases {
		result := frequentElements(c.arr, c.k)
		assert.ElementsMatch(t, c.expectedResult, result)
	}
}
