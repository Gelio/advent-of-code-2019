package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameA(t *testing.T) {
	p1 := player{
		ID:   1,
		deck: newDeck([]int{9, 2, 6, 3, 1}),
	}

	p2 := player{
		ID:   2,
		deck: newDeck([]int{5, 8, 4, 7, 10}),
	}

	winner := playGameA(&p1, &p2)

	assert.Equal(t, &p2, winner, "wrong winner")

	winningCards := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	assert.Equal(t, winningCards, winner.Cards(), "wrong winning cards")
}
