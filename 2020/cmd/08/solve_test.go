package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var input = strings.Split(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`, "\n")

func TestSolveA(t *testing.T) {
	res, err := SolveA(input)
	require.NoError(t, err)

	assert.Equal(t, 5, res)
}

func TestSolveB(t *testing.T) {
	res, err := SolveB(input)
	require.NoError(t, err)

	assert.Equal(t, 8, res)
}
