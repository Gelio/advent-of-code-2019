package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	p1 := player{
		ID:   1,
		deck: newDeck([]int{9, 2, 6, 3, 1}),
	}

	p2 := player{
		ID:   2,
		deck: newDeck([]int{5, 8, 4, 7, 10}),
	}

	assert.Equal(t, 306, solveA(&p1, &p2), "invalid result")
}
