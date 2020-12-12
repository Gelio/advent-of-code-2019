package waypoint

import "aoc-2020/cmd/12/ship"

type Waypoint struct {
	// Position relative to the ship
	X, Y int
	ship.Ship
}

func New() Waypoint {
	return Waypoint{
		Ship: ship.New(),
		X:    10,
		Y:    1,
	}
}
