package internal

const (
	emptySeat    = 'L'
	occupiedSeat = '#'
	floor        = '.'
)

// SeatMap is a 2-dimentional array of seats
type SeatMap [][]rune

type position struct {
	x, y int
}

// ParseSeatMap parses lines into a seat map
func ParseSeatMap(lines []string) SeatMap {
	var m SeatMap

	for _, l := range lines {
		m = append(m, []rune(l))
	}

	return m
}

func (m SeatMap) get(x, y int) rune {
	return m[y][x]
}

func (m SeatMap) set(x, y int, v rune) {
	m[y][x] = v
}

func (m SeatMap) forEach(f func(x, y int, v rune)) {
	for y, row := range m {
		for x, v := range row {
			f(x, y, v)
		}
	}
}

func (m SeatMap) isWithin(p position) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(m) && p.x < len(m[0])
}

// Simulate runs the seat-changing simulation described in the puzzle
func (m SeatMap) Simulate(maxDist, minOccupiedNeighborsToLeave int) (simulationRuns int) {
	dirty := true

	ns := getNeighboringSeats(m, maxDist)
	occupiedNeighborsCount := make(map[position]int)

	for dirty {
		dirty = false
		simulationRuns++

		for pos, neighbors := range ns {
			if m.get(pos.x, pos.y) == occupiedSeat {
				for _, n := range neighbors {
					occupiedNeighborsCount[n]++
				}
			}
		}

		for pos := range ns {
			nc := occupiedNeighborsCount[pos]
			occupiedNeighborsCount[pos] = 0

			switch m.get(pos.x, pos.y) {
			case floor:
				continue

			case emptySeat:
				if nc == 0 {
					dirty = true
					m.set(pos.x, pos.y, occupiedSeat)
				}

			case occupiedSeat:
				if nc >= minOccupiedNeighborsToLeave {
					dirty = true
					m.set(pos.x, pos.y, emptySeat)
				}
			}
		}
	}

	return
}

// CountOccupiedSeats counts the number of occupied seats
func (m SeatMap) CountOccupiedSeats() int {
	occupiedSeats := 0

	m.forEach(func(_, _ int, v rune) {
		if v == occupiedSeat {
			occupiedSeats++
		}
	})

	return occupiedSeats
}
