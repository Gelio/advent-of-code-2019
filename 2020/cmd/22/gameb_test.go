package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var gameBCases = []struct {
	name           string
	getDecks       func() (deck, deck, error)
	expectedWinner int
	winningCards   []int
}{
	{
		name: "small input",
		getDecks: func() (deck, deck, error) {
			d1 := newDeck([]int{9, 2, 6, 3, 1})
			d2 := newDeck([]int{5, 8, 4, 7, 10})

			return d1, d2, nil
		},
		expectedWinner: 2,
		winningCards:   []int{7, 5, 6, 2, 4, 1, 10, 8, 9, 3},
	},
	{
		name:           "full input",
		getDecks:       getPlayerDecksFromInput,
		expectedWinner: 2,
		winningCards:   []int{1, 6, 49, 29, 30, 18, 32, 25, 35, 20, 21, 3, 46, 43, 45, 8, 16, 7, 42, 24, 39, 19, 48, 47, 5, 4, 44, 15, 22, 2, 38, 31, 34, 28, 41, 17, 37, 12, 27, 14, 26, 10, 50, 23, 36, 11, 40, 13, 33, 9},
	},
}

func TestGameB(t *testing.T) {
	for _, tt := range gameBCases {
		t.Run(fmt.Sprintf("%s", tt.name), func(t *testing.T) {
			d1, d2, err := tt.getDecks()
			require.NoError(t, err, "cannot get decks")

			winner := playGameB(&d1, &d2)
			if tt.expectedWinner == 1 {
				assert.Equal(t, &d1, winner, "wrong winner")
			} else {
				assert.Equal(t, &d2, winner, "wrong winner")
			}

			assert.Equal(t, tt.winningCards, winner.Cards(), "wrong winning cards")
		})
	}
}

func BenchmarkGameB(b *testing.B) {
	for _, tt := range gameBCases {
		b.Run(fmt.Sprintf("%s", tt.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				d1, d2, err := tt.getDecks()
				require.NoError(b, err, "cannot get decks")

				playGameB(&d1, &d2)
			}
		})
	}
}
