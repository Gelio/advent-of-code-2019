package main

import "fmt"

type move int

const (
	NW move = iota
	NE
	E
	SE
	SW
	W
)

type position struct {
	x, y int
}

func (p position) getNeighboringTilesPositions() []position {
	return []position{
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x - 1, p.y + 1},
		{p.x - 1, p.y},
	}
}

func getPosition(moves []move) position {
	var p position
	for _, m := range moves {
		switch m {
		case NW:
			p.y--
		case NE:
			p.y--
			p.x++
		case E:
			p.x++
		case SE:
			p.y++
		case SW:
			p.y++
			p.x--
		case W:
			p.x--
		}
	}

	return p
}

func parseMoves(line string) ([]move, error) {
	var moves []move

	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case 'e':
			moves = append(moves, E)
		case 'w':
			moves = append(moves, W)
		case 's':
			if peek(line, 'w', i) {
				moves = append(moves, SW)
				i++
			} else if peek(line, 'e', i) {
				moves = append(moves, SE)
				i++
			} else {
				return nil, fmt.Errorf("e or w expected after s at position %d", i+1)
			}
		case 'n':
			if peek(line, 'w', i) {
				moves = append(moves, NW)
				i++
			} else if peek(line, 'e', i) {
				moves = append(moves, NE)
				i++
			} else {
				return nil, fmt.Errorf("e or w expected after n at position %d", i+1)
			}
		default:
			return nil, fmt.Errorf("unexpected character %v at position %d", c, i+1)
		}
	}

	return moves, nil
}

func peek(line string, c byte, i int) bool {
	if i+1 >= len(line) {
		return false
	}

	return line[i+1] == c
}
