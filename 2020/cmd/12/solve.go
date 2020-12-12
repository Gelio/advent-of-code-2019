package main

import (
	"aoc-2020/cmd/12/instructions"
	"aoc-2020/cmd/12/ship"
	"math"
)

func solveA(input []string) (int, error) {
	var instrs []instructions.Instruction

	for _, l := range input {
		instr, err := instructions.Parse(l)
		if err != nil {
			return 0, err
		}

		instrs = append(instrs, instr)
	}

	s := ship.New()

	for _, instr := range instrs {
		err := instr.Exec(&s)

		if err != nil {
			return 0, err
		}
	}

	return getManhattanDistance(s), nil
}

func getManhattanDistance(s ship.Ship) int {
	return int(math.Abs(float64(s.X)) + math.Abs(float64(s.Y)))
}
