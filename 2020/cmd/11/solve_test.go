package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = strings.Split(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, "\n")

func TestSolveA(t *testing.T) {
	res := solveA(input)

	assert.Equal(t, 37, res)
}

func TestSolveB(t *testing.T) {
	res := solveB(input)

	assert.Equal(t, 26, res)
}
