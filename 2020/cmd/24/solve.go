package main

import "fmt"

func solveA(lines []string) (int, error) {
	blackTiles, err := getInitialBlackTiles(lines)

	return len(blackTiles), err
}

func solveB(lines []string, iterations int) (int, error) {
	blackTiles, err := getInitialBlackTiles(lines)
	if err != nil {
		return 0, fmt.Errorf("cannot get initial black tiles: %w", err)
	}

	for i := 0; i < iterations; i++ {
		neighboringBlackTiles := make(map[position]int)
		for pos := range blackTiles {
			if _, ok := neighboringBlackTiles[pos]; !ok {
				// Make sure any existing black tile is processed
				neighboringBlackTiles[pos] = 0
			}

			for _, neighborPos := range pos.getNeighboringTilesPositions() {
				neighboringBlackTiles[neighborPos]++
			}
		}

		for pos, blackTileNeighbors := range neighboringBlackTiles {
			blackTile := blackTiles[pos]
			if blackTile {
				if blackTileNeighbors == 0 || blackTileNeighbors > 2 {
					delete(blackTiles, pos)
				}
			} else {
				if blackTileNeighbors == 2 {
					blackTiles[pos] = true
				}
			}
		}
	}

	return len(blackTiles), nil
}

func getInitialBlackTiles(lines []string) (map[position]bool, error) {
	blackTiles := make(map[position]bool)

	for _, line := range lines {
		moves, err := parseMoves(line)
		if err != nil {
			return nil, fmt.Errorf("cannot parse line %q: %w", line, err)
		}

		pos := getPosition(moves)
		if blackTiles[pos] {
			delete(blackTiles, pos)
		} else {
			blackTiles[pos] = true
		}
	}

	return blackTiles, nil
}
