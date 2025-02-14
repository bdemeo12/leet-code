package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	type params struct {
		inputString  string
		outputString string
	}

	cases := []params{
		{
			inputString:  "leetcode",
			outputString: "leotcede",
		},
		{
			inputString:  "IceCreAm",
			outputString: "AceCreIm",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.outputString, reverseVowels(c.inputString))
	}
}
