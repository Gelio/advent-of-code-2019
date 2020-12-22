package main

import (
	"aoc-2020/internal/parse"
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"fmt"
)

func main() {
	input, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	playerDefinitions := testcases.SplitTestCaseLines(input)

	player1Cards, err := parse.Ints(playerDefinitions[0][1:])
	if err != nil {
		fmt.Println("Error parsing player 1 cards:", err)
		return
	}
	d1 := newDeck(player1Cards)

	player2Cards, err := parse.Ints(playerDefinitions[1][1:])
	if err != nil {
		fmt.Println("Error parsing player 2 cards:", err)
		return
	}
	d2 := newDeck(player2Cards)

	fmt.Println("Result A:", solveA(d1, d2))
	fmt.Println("Result B:", solveB(d1, d2))
}
