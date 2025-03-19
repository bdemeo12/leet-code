package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	type params struct {
		input          []int
		expectedOutput bool
	}

	cases := []params{
		{
			input:          []int{20, 100, 10, 12, 5, 13},
			expectedOutput: true,
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: true,
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedOutput: false,
		},
		{
			input:          []int{2, 1, 5, 0, 4, 6},
			expectedOutput: true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, increasingTriplet(c.input))
	}
}
