package main

const maxDeckSize = 50

type gameBDeckCards [maxDeckSize]int

func getGameBDeckCards(d deck) gameBDeckCards {
	var cards gameBDeckCards

	c := d.Top
	i := 0
	for c != nil {
		cards[i] = c.Val
		i++
		c = c.Next
	}

	return cards
}

func playGameB(p1, p2 *player) (winner *player) {
	encounteredDecksP1 := make(map[gameBDeckCards]bool)
	encounteredDecksP2 := make(map[gameBDeckCards]bool)

	for p1.deck.Length > 0 && p2.deck.Length > 0 {
		// Deal with cache
		p1AllCards := getGameBDeckCards(p1.deck)
		p2AllCards := getGameBDeckCards(p2.deck)
		if encounteredDecksP1[p1AllCards] && encounteredDecksP2[p2AllCards] {
			return p1
		}

		encounteredDecksP1[p1AllCards] = true
		encounteredDecksP2[p2AllCards] = true

		// Drawing cards and determining which game to play
		c1 := p1.deck.PopCard()
		c2 := p2.deck.PopCard()

		var winnerDeck *deck
		var betterCard, worseCard *card

		// Check deck lengths and possibly recurse
		if p1.deck.Length >= c1.Val && p2.deck.Length >= c2.Val {
			p1Clone, p2Clone := p1.CloneWithLength(c1.Val), p2.CloneWithLength(c2.Val)
			nextRoundWinner := playGameB(&p1Clone, &p2Clone)

			if nextRoundWinner == &p1Clone {
				winnerDeck = &p1.deck
				betterCard = c1
				worseCard = c2
			} else {
				winnerDeck = &p2.deck
				betterCard = c2
				worseCard = c1
			}
		} else {
			// Regular round
			if c1.Val > c2.Val {
				winnerDeck = &p1.deck
				betterCard = c1
				worseCard = c2
			} else {
				winnerDeck = &p2.deck
				betterCard = c2
				worseCard = c1
			}
		}

		winnerDeck.AddCard(betterCard)
		winnerDeck.AddCard(worseCard)
	}

	if p1.deck.Length > 0 {
		return p1
	}

	return p2
}
