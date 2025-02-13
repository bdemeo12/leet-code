package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerWithMostWater(t *testing.T) {
	type params struct {
		input          []int
		expectedOutput int
	}

	cases := []params{
		{
			input:          []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expectedOutput: 49,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, maxArea(c.input))
	}
}
