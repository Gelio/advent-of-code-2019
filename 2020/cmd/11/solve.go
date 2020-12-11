package main

func solveA(lines []string) int {
	m := parseSeatMap(lines)

	m.simulate()

	return m.countOccupiedSeats()
}

type seatMap [][]rune

const (
	emptySeat    = 'L'
	occupiedSeat = '#'
	floor        = '.'
)

func parseSeatMap(lines []string) seatMap {
	var m seatMap

	for _, l := range lines {
		m = append(m, []rune(l))
	}

	return m
}

func (m seatMap) simulate() (simulationRuns int) {
	dirty := true

	neighborsCount := make([][]int, len(m))
	for i := range m {
		neighborsCount[i] = make([]int, len(m[i]))
	}

	for dirty {
		dirty = false
		simulationRuns++

		for y, row := range m {
			for x, v := range row {
				if v == occupiedSeat {
					markSeatOccupied(neighborsCount, x, y)
				}
			}
		}

		for y, row := range m {
			for x, v := range row {
				nc := neighborsCount[y][x]
				neighborsCount[y][x] = 0

				switch v {
				case floor:
					continue

				case emptySeat:
					if nc == 0 {
						dirty = true
						row[x] = occupiedSeat
					}

				case occupiedSeat:
					if nc >= 4 {
						dirty = true
						row[x] = emptySeat
					}
				}
			}
		}
	}

	return
}

func markSeatOccupied(neighborsCount [][]int, x, y int) {
	for i := max(0, y-1); i <= min(len(neighborsCount)-1, y+1); i++ {
		for j := max(0, x-1); j <= min(len(neighborsCount[i])-1, x+1); j++ {
			if i == y && j == x {
				continue
			}

			neighborsCount[i][j]++
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (m seatMap) countOccupiedSeats() int {
	occupiedSeats := 0

	for _, row := range m {
		for _, v := range row {
			if v == occupiedSeat {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}
