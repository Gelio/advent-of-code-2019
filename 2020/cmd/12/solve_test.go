package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var input = strings.Split(`F10
N3
F7
R90
F11`, "\n")

func TestSolveA(t *testing.T) {
	res, err := solveA(input)

	require.NoError(t, err)

	assert.Equal(t, 25, res)
}

func TestSolveB(t *testing.T) {
	res, err := solveB(input)

	require.NoError(t, err)

	assert.Equal(t, 286, res)
}
