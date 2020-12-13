package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBusIDs(t *testing.T) {
	busIDs, err := parseBusIDs("7,13,x,x,59,x,31,19")

	require.NoError(t, err)

	assert.Equal(t, busIDs, []int{7, 13, 59, 31, 19})
}

func TestSolveA(t *testing.T) {
	timestamp := 939
	busIDs := []int{7, 13, 59, 31, 19}

	result := solveA(timestamp, busIDs)

	assert.Equal(t, 295, result)
}
