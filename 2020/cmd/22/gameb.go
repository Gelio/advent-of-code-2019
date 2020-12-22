package main

func playGameB(p1, p2 *player) (winner *player) {
	for p1.deck.Length > 0 && p2.deck.Length > 0 {
		c1 := p1.deck.PopCard()
		c2 := p2.deck.PopCard()

		var winnerDeck *deck
		var betterCard, worseCard *card
		if c1.Val > c2.Val {
			winnerDeck = &p1.deck
			betterCard = c1
			worseCard = c2
		} else {
			winnerDeck = &p2.deck
			betterCard = c2
			worseCard = c1
		}

		winnerDeck.AddCard(betterCard)
		winnerDeck.AddCard(worseCard)
	}

	if p1.deck.Length > 0 {
		return p1
	}

	return p2
}
