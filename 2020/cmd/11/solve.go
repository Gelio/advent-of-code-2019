package main

import "aoc-2020/cmd/11/internal"

func solveA(lines []string) int {
	m := internal.ParseSeatMap(lines)

	m.Simulate(1, 4)

	return m.CountOccupiedSeats()
}

func solveB(lines []string) int {
	m := internal.ParseSeatMap(lines)

	m.Simulate(1e5, 5)

	return m.CountOccupiedSeats()
}
