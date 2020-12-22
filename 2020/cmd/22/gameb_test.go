package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameB(t *testing.T) {
	d1 := newDeck([]int{9, 2, 6, 3, 1})
	d2 := newDeck([]int{5, 8, 4, 7, 10})

	winner := playGameB(&d1, &d2)

	assert.Equal(t, &d2, winner, "wrong winner")

	winningCards := []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}
	assert.Equal(t, winningCards, winner.Cards(), "wrong winning cards")
}

func BenchmarkGameB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d1 := newDeck([]int{9, 2, 6, 3, 1})
		d2 := newDeck([]int{5, 8, 4, 7, 10})

		playGameB(&d1, &d2)
	}
}
