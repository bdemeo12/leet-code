package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreatestCommonDivisor(t *testing.T) {
	type gcdParams struct {
		str1           string
		str2           string
		expectedOutput string
	}

	cases := []gcdParams{
		{
			str1:           "ABCABC",
			str2:           "ABC",
			expectedOutput: "ABC",
		},
		{
			str1:           "ABABAB",
			str2:           "ABAB",
			expectedOutput: "AB",
		},
		{
			str1:           "LEET",
			str2:           "CODE",
			expectedOutput: "",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedOutput, gcdOfStrings(c.str1, c.str2))
	}
}
