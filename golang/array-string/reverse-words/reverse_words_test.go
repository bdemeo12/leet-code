package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	type params struct {
		s              string
		expectedResult string
	}

	cases := []params{
		{
			s:              "hello world",
			expectedResult: "world hello",
		},
		{
			s:              "the sky is blue",
			expectedResult: "blue is sky the",
		},
		// {
		// 	s:              "a good   example",
		// 	expectedResult: "example good a",
		// },
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedResult, reverseWords(c.s))
	}
}
