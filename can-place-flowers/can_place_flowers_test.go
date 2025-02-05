package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	type params struct {
		flowerbed      []int
		n              int
		expectedResult bool
	}

	cases := []params{
		{
			flowerbed:      []int{1, 0, 0, 0, 1},
			n:              1,
			expectedResult: true,
		},
		{
			flowerbed:      []int{1, 0, 0, 0, 1},
			n:              2,
			expectedResult: false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedResult, canPlaceFlowers(c.flowerbed, c.n))
	}
}
