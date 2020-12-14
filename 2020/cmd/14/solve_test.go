package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	input := strings.Split(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`, "\n")

	res, err := solveA(input)

	require.NoError(t, err)

	assert.Equal(t, 165, res)
}

func TestSolveB(t *testing.T) {
	input := strings.Split(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, "\n")

	res, err := solveB(input)

	require.NoError(t, err)

	assert.Equal(t, 208, res)
}
