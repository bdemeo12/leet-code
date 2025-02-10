package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {
	type params struct {
		inputArr       []int
		expectedOutput []int
	}

	cases := []params{
		{
			inputArr:       []int{0, 1, 0, 3, 12},
			expectedOutput: []int{1, 3, 12, 0, 0},
		},
		{
			inputArr:       []int{0},
			expectedOutput: []int{0},
		},
		{
			inputArr:       []int{0, 1},
			expectedOutput: []int{1, 0},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, moveZeros(c.inputArr))
	}
}
