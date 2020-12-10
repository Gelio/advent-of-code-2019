package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var input = strings.Split(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`, "\n")

func TestSolveA(t *testing.T) {
	nums, err := parseNums(input)
	require.NoError(t, err)

	result, err := solveA(nums, 5)

	require.NoError(t, err)
	assert.Equal(t, 127, result)
}

func TestSolveB(t *testing.T) {
	nums, err := parseNums(input)
	require.NoError(t, err)

	min, max, err := solveB(nums, 127)

	require.NoError(t, err)
	assert.Equal(t, 15, min)
	assert.Equal(t, 47, max)
}
