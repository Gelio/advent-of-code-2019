package main

// A trie tree for storing decks encountered during a round
// Each node represents a card at particular distance from the top of the deck.
// Edges point to next cards in the deck
type gameBDecks map[int]gameBDecks

func (decks gameBDecks) AddDeck(d *deck) (existingDeck bool) {
	return decks.addCards(d.Top)
}

func (decks gameBDecks) addCards(c *card) (existingDeck bool) {
	if c == nil {
		// Use -1 to mark that some deck ended here
		if _, ok := decks[-1]; ok {
			return true
		}

		decks[-1] = nil
		return false
	}

	nextTrieNode, ok := decks[c.Val]
	if ok {
		return nextTrieNode.addCards(c.Next)
	}

	decks[c.Val] = make(gameBDecks)
	decks[c.Val].addCards(c.Next)
	return false
}

func playGameB(d1, d2 *deck) (winner *deck) {
	encounteredDecksP1 := make(gameBDecks)
	encounteredDecksP2 := make(gameBDecks)

	for d1.Length > 0 && d2.Length > 0 {
		// Deal with cache
		if encounteredDecksP1.AddDeck(d1) && encounteredDecksP2.AddDeck(d2) {
			return d1
		}

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
