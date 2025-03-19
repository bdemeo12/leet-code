package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "olleh", reverseString("hello"))
}
