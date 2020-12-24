package main

import "fmt"

func solveA(lines []string) (int, error) {
	blackTiles := make(map[position]bool)

	for _, line := range lines {
		moves, err := parseMoves(line)
		if err != nil {
			return 0, fmt.Errorf("cannot parse line %q: %w", line, err)
		}

		pos := getPosition(moves)
		if blackTiles[pos] {
			delete(blackTiles, pos)
		} else {
			blackTiles[pos] = true
		}
	}

	return len(blackTiles), nil
}
