package main

func playGameA(d1, d2 *deck) (winner *deck) {
	for d1.Length > 0 && d2.Length > 0 {
		c1 := d1.PopCard()
		c2 := d2.PopCard()

		var winnerDeck *deck
		var betterCard, worseCard *card
		if c1.Val > c2.Val {
			winnerDeck = d1
			betterCard = c1
			worseCard = c2
		} else {
			winnerDeck = d2
			betterCard = c2
			worseCard = c1
		}

		winnerDeck.AddCard(betterCard)
		winnerDeck.AddCard(worseCard)
	}

	if d1.Length > 0 {
		return d1
	}

	return d2
}
