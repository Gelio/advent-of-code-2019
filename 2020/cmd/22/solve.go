package main

func solveA(p1, p2 *player) int {
	winner := playGame(p1, p2)

	winnerCards := winner.Cards()

	result := 0

	for i, c := range winnerCards {
		result += (winner.deck.Length - i) * c
	}

	return result
}
