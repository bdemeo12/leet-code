package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequenece(t *testing.T) {
	type params struct {
		s      string
		t      string
		output bool
	}

	cases := []params{
		{
			s:      "abc",
			t:      "ahbgdc",
			output: true,
		},
		{
			s:      "acx",
			t:      "ahbgdc",
			output: false,
		},
		{
			s:      "acb",
			t:      "ahbgdc",
			output: false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.output, isSubsequence(c.s, c.t))
	}
}
