package main

type deck struct {
	Top, Bottom *card
	Length      int
}

func newDeck(cards []int) deck {
	var d deck

	for _, c := range cards {
		d.AddCard(&card{Val: c})
	}

	return d
}

type card struct {
	Val  int
	Next *card
}

func (d *deck) AddCard(c *card) {
	if d.Length > 0 {
		d.Bottom.Next = c
	} else {
		d.Top = c
	}
	c.Next = nil
	d.Bottom = c
	d.Length++
}

func (d *deck) PopCard() *card {
	c := d.Top
	d.Top = d.Top.Next
	c.Next = nil
	d.Length--

	if d.Length == 0 {
		d.Bottom = nil
	}

	return c
}

// NOTE: probably unnecessary
func (d deck) Cards() []int {
	var cards []int

	c := d.Top
	for c != nil {
		cards = append(cards, c.Val)
		c = c.Next
	}

	return cards
}

func (d deck) Score() int {
	// NOTE: could be optimized by iterating over cards directly, instead of calling .Cards
	winnerCards := d.Cards()

	result := 0

	for i, c := range winnerCards {
		result += (d.Length - i) * c
	}

	return result
}

func (d deck) Clone() deck {
	// NOTE: could be optimized by iterating over cards directly, instead of calling .Cards
	cards := d.Cards()

	return newDeck(cards)
}

func (d deck) CloneWithLength(length int) deck {
	// NOTE: could be optimized by iterating over cards directly, instead of calling .Cards
	cards := d.Cards()[0:length]

	return newDeck(cards)
}
