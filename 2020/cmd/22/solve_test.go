package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	d1 := newDeck([]int{9, 2, 6, 3, 1})
	d2 := newDeck([]int{5, 8, 4, 7, 10})

	assert.Equal(t, 306, solveA(d1, d2), "invalid result")
}

func TestSolveB(t *testing.T) {
	d1 := newDeck([]int{9, 2, 6, 3, 1})
	d2 := newDeck([]int{5, 8, 4, 7, 10})

	assert.Equal(t, 291, solveB(d1, d2), "invalid result")
}
