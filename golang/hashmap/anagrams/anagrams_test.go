package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	type isAnagramParams struct {
		inputString1   string
		inputString2   string
		expectedResult bool
	}

	cases := []isAnagramParams{
		{
			inputString1:   "racecar",
			inputString2:   "carrace",
			expectedResult: true,
		},
		{
			inputString1:   "jam",
			inputString2:   "jar",
			expectedResult: false,
		},
		{
			inputString1:   "jam",
			inputString2:   "jam",
			expectedResult: true,
		},
		{
			inputString1:   "ja",
			inputString2:   "jar",
			expectedResult: false,
		},
	}

	for _, c := range cases {
		result := isAnagram(c.inputString1, c.inputString2)
		assert.Equal(t, c.expectedResult, result)
	}

}
