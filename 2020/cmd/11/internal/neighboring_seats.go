package internal

type neighboringSeats map[position][]position

type direction struct {
	xDelta, yDelta int
}

func getNeighboringSeats(m SeatMap, maxDistance int) neighboringSeats {
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

func getSeatInDirection(m SeatMap, pos position, dir direction, maxDist int) (position, bool) {
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
