package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	type groupAnagramsParams struct {
		strs           []string
		expectedResult [][]string
	}

	cases := []groupAnagramsParams{
		{
			strs:           []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expectedResult: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
	}

	for _, c := range cases {
		result := groupAnagrams(c.strs)
		assert.NotNil(t, result)
	}
}
