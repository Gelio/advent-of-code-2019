package main

func solveA(lines []string) int {
	m := parseSeatMap(lines)

	m.simulate(1, 4)

	return m.countOccupiedSeats()
}

func solveB(lines []string) int {
	m := parseSeatMap(lines)

	m.simulate(1e5, 5)

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

func (m seatMap) get(x, y int) rune {
	return m[y][x]
}

func (m seatMap) set(x, y int, v rune) {
	m[y][x] = v
}

func (m seatMap) forEach(f func(x, y int, v rune)) {
	for y, row := range m {
		for x, v := range row {
			f(x, y, v)
		}
	}
}

func (m seatMap) isWithin(p position) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(m) && p.x < len(m[0])
}

func (m seatMap) simulate(maxDist, minOccupiedNeighborsToLeave int) (simulationRuns int) {
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

func (m seatMap) countOccupiedSeats() int {
	occupiedSeats := 0

	m.forEach(func(_, _ int, v rune) {
		if v == occupiedSeat {
			occupiedSeats++
		}
	})

	return occupiedSeats
}

type position struct {
	x, y int
}

type neighboringSeats map[position][]position

type direction struct {
	xDelta, yDelta int
}

func getNeighboringSeats(m seatMap, maxDistance int) neighboringSeats {
	directions := []direction{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	ns := make(neighboringSeats)

	m.forEach(func(x, y int, v rune) {
		if v != emptySeat {
			return
		}

		pos := position{x, y}

		ns[pos] = nil

		for _, d := range directions {
			if p, found := getSeatInDirection(m, pos, d, maxDistance); found {
				ns[pos] = append(ns[pos], p)
			}
		}
	})

	return ns
}

func getSeatInDirection(m seatMap, pos position, dir direction, maxDist int) (position, bool) {
	for i := 0; i < maxDist; i++ {
		pos.x += dir.xDelta
		pos.y += dir.yDelta

		if !m.isWithin(pos) {
			break
		}

		v := m.get(pos.x, pos.y)

		if v != floor {
			return position{pos.x, pos.y}, true
		}
	}

	return position{}, false
}
