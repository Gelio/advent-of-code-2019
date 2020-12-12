package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	input := strings.Split(`F10
N3
F7
R90
F11`, "\n")

	res, err := solveA(input)

	require.NoError(t, err)

	assert.Equal(t, 25, res)
}
