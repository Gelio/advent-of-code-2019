package main

const maxDeckSize = 50

type gameBDeckCards [maxDeckSize]int

func getGameBDeckCards(d *deck) gameBDeckCards {
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

func playGameB(d1, d2 *deck) (winner *deck) {
	encounteredDecksP1 := make(map[gameBDeckCards]bool)
	encounteredDecksP2 := make(map[gameBDeckCards]bool)

	for d1.Length > 0 && d2.Length > 0 {
		// Deal with cache
		p1AllCards := getGameBDeckCards(d1)
		p2AllCards := getGameBDeckCards(d2)
		if encounteredDecksP1[p1AllCards] && encounteredDecksP2[p2AllCards] {
			return d1
		}

		encounteredDecksP1[p1AllCards] = true
		encounteredDecksP2[p2AllCards] = true

		// Drawing cards and determining which game to play
		c1 := d1.PopCard()
		c2 := d2.PopCard()

		var winnerDeck *deck
		var betterCard, worseCard *card

		// Check deck lengths and possibly recurse
		if d1.Length >= c1.Val && d2.Length >= c2.Val {
			p1Clone, p2Clone := d1.CloneWithLength(c1.Val), d2.CloneWithLength(c2.Val)
			nextRoundWinner := playGameB(&p1Clone, &p2Clone)

			if nextRoundWinner == &p1Clone {
				winnerDeck = d1
				betterCard = c1
				worseCard = c2
			} else {
				winnerDeck = d2
				betterCard = c2
				worseCard = c1
			}
		} else {
			// Regular round
			if c1.Val > c2.Val {
				winnerDeck = d1
				betterCard = c1
				worseCard = c2
			} else {
				winnerDeck = d2
				betterCard = c2
				worseCard = c1
			}
		}

		winnerDeck.AddCard(betterCard)
		winnerDeck.AddCard(worseCard)
	}

	if d1.Length > 0 {
		return d1
	}

	return d2
}
