package main

import (
	"aoc-2020/internal/parse"
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"fmt"
)

func solveA(d1, d2 deck) int {
	p1Clone, p2Clone := d1.Clone(), d2.Clone()
	winner := playGameA(&p1Clone, &p2Clone)

	return winner.Score()
}

func solveB(d1, d2 deck) int {
	p1Clone, p2Clone := d1.Clone(), d2.Clone()
	winner := playGameB(&p1Clone, &p2Clone)

	return winner.Score()
}

func getPlayerDecksFromInput() (deck, deck, error) {
	input, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		return deck{}, deck{}, fmt.Errorf("cannot read input: %w", err)
	}

	playerDefinitions := testcases.SplitTestCaseLines(input)

	player1Cards, err := parse.Ints(playerDefinitions[0][1:])
	if err != nil {
		return deck{}, deck{}, fmt.Errorf("cannot parse player 1 cards: %w", err)
	}
	d1 := newDeck(player1Cards)

	player2Cards, err := parse.Ints(playerDefinitions[1][1:])
	if err != nil {
		return deck{}, deck{}, fmt.Errorf("cannot parse player 2 cards: %w", err)
	}
	d2 := newDeck(player2Cards)

	return d1, d2, nil
}
