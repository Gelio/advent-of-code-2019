package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	res, err := solveA(strings.Split(`.#.
..#
###`, "\n"))

	require.NoError(t, err, "error when parsing input")

	assert.Equal(t, 112, res)
}
