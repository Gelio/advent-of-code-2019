package main

import "fmt"

const (
	inactiveCube = '.'
	activeCube   = '#'
)

type position struct {
	X, Y, Z int
}

func solveA(lines []string) (int, error) {
	activeCubePositions, err := getActiveCubePositions(lines)
	if err != nil {
		return 0, err
	}

	const cycles = 6

	for i := 0; i < cycles; i++ {
		neighboringActiveCubeCounts := make(map[position]int)

		for p := range activeCubePositions {
			for _, n := range getNeighboringPositions(p) {
				neighboringActiveCubeCounts[n]++
			}
		}

		nextCycleActiveCubePositions := make(map[position]bool)

		for p, count := range neighboringActiveCubeCounts {
			if activeCubePositions[p] {
				if count == 2 || count == 3 {
					nextCycleActiveCubePositions[p] = true
				}
			} else {
				if count == 3 {
					nextCycleActiveCubePositions[p] = true
				}
			}
		}

		activeCubePositions = nextCycleActiveCubePositions
	}

	return len(activeCubePositions), nil
}

func getActiveCubePositions(lines []string) (map[position]bool, error) {
	p := make(map[position]bool)

	for y, line := range lines {
		for x, c := range line {
			switch c {
			case inactiveCube:

			case activeCube:
				p[position{X: x, Y: y}] = true

			default:
				return nil, fmt.Errorf("unknown character %c at position %d in line %d", c, x, y)
			}
		}
	}

	return p, nil
}

func getNeighboringPositions(p position) []position {
	var neighbors []position

	for x := p.X - 1; x <= p.X+1; x++ {
		for y := p.Y - 1; y <= p.Y+1; y++ {
			for z := p.Z - 1; z <= p.Z+1; z++ {
				if x == p.X && y == p.Y && z == p.Z {
					continue
				}

				neighbors = append(neighbors, position{x, y, z})
			}
		}
	}

	return neighbors
}
