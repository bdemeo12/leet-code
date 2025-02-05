package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeStringsALernately(t *testing.T) {
	type mergeStringsParams struct {
		word1          string
		word2          string
		expectedOutput string
	}

	cases := []mergeStringsParams{
		{
			word1:          "abc",
			word2:          "pqr",
			expectedOutput: "apbqcr",
		},
		{
			word1:          "ab",
			word2:          "pqrs",
			expectedOutput: "apbqrs",
		},
		{
			word1:          "abcd",
			word2:          "pq",
			expectedOutput: "apbqcd",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, mergeAlternately(c.word1, c.word2))
	}
}
