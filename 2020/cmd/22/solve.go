package main

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
