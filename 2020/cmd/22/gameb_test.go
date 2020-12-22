package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameB(t *testing.T) {
	p1 := player{
		ID:   1,
		deck: newDeck([]int{9, 2, 6, 3, 1}),
	}

	p2 := player{
		ID:   2,
		deck: newDeck([]int{5, 8, 4, 7, 10}),
	}

	winner := playGameB(&p1, &p2)

	assert.Equal(t, &p2, winner, "wrong winner")

	winningCards := []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3}
	assert.Equal(t, winningCards, winner.Cards(), "wrong winning cards")
}

func BenchmarkGameB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p1 := player{
			ID:   1,
			deck: newDeck([]int{9, 2, 6, 3, 1}),
		}

		p2 := player{
			ID:   2,
			deck: newDeck([]int{5, 8, 4, 7, 10}),
		}

		playGameB(&p1, &p2)
	}
}
