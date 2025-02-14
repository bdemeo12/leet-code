package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestValue(t *testing.T) {
	m := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 92,
	}
	assert.Equal(t, 92, largestValue(m))
}
