package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	type params struct {
		chars          []byte
		expectedOutput int
	}

	cases := []params{
		{
			chars:          []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expectedOutput: 6,
		},
		{
			chars:          []byte{'a'},
			expectedOutput: 1,
		},
		{
			chars:          []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			expectedOutput: 4,
		},
	}

	for _, c := range cases {
		output := compress(c.chars)
		assert.Equal(t, c.expectedOutput, output)
	}
}
