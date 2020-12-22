package main

func solveA(p1, p2 player) int {
	p1Clone, p2Clone := p1.Clone(), p2.Clone()
	winner := playGameA(&p1Clone, &p2Clone)

	return winner.Score()
}

func solveB(p1, p2 player) int {
	p1Clone, p2Clone := p1.Clone(), p2.Clone()
	winner := playGameB(&p1Clone, &p2Clone)

	return winner.Score()
}
