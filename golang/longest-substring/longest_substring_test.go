package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	type params struct {
		input          string
		k              int
		expectedOutput int
	}

	cases := []params{
		{
			input:          "ababbc",
			k:              2,
			expectedOutput: 2,
		},
		{
			input:          "aaabb",
			k:              3,
			expectedOutput: 3,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, longestSubstring(c.input, c.k))
	}
}
