package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameA(t *testing.T) {
	d1 := newDeck([]int{9, 2, 6, 3, 1})
	d2 := newDeck([]int{5, 8, 4, 7, 10})

	winner := playGameA(&d1, &d2)

	assert.Equal(t, &d2, winner, "wrong winner")

	winningCards := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	assert.Equal(t, winningCards, winner.Cards(), "wrong winning cards")
}
