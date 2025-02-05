package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreatestNumberOfCandies(t *testing.T) {
	type params struct {
		candies        []int
		extraCandies   int
		expectedOutput []bool
	}

	cases := []params{
		{
			candies:        []int{4, 2, 1, 1, 2},
			extraCandies:   1,
			expectedOutput: []bool{true, false, false, false, false},
		},
		{
			candies:        []int{12, 1, 12},
			extraCandies:   10,
			expectedOutput: []bool{true, false, true},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, kidsWithCandies(c.candies, c.extraCandies))
	}

}
