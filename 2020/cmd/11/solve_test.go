package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	input := strings.Split(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n")

	res := solveA(input)

	assert.Equal(t, 37, res)
}
