package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var input = strings.Split(`.#.
..#
###`, "\n")

func TestSolveA(t *testing.T) {
	res, err := solve(input, getNeighboringPositions3D)

	require.NoError(t, err, "error when parsing input")

	assert.Equal(t, 112, res)
}

func TestSolveB(t *testing.T) {
	res, err := solve(input, getNeighboringPositions4D)

	require.NoError(t, err, "error when parsing input")

	assert.Equal(t, 848, res)
}
