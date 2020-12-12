package main

import (
	"aoc-2020/cmd/12/instructions"
	"aoc-2020/cmd/12/ship"
	"aoc-2020/cmd/12/waypoint"
	"math"
)

func solveA(input []string) (int, error) {
	instrs, err := parseInstrs(input)
	if err != nil {
		return 0, err
	}

	s := ship.New()

	for _, instr := range instrs {
		err := instr.ExecShip(&s)

		if err != nil {
			return 0, err
		}
	}

	return getManhattanDistance(s), nil
}

func solveB(input []string) (int, error) {
	instrs, err := parseInstrs(input)
	if err != nil {
		return 0, err
	}

	w := waypoint.New()

	for _, instr := range instrs {
		err := instr.ExecWaypoint(&w)

		if err != nil {
			return 0, err
		}
	}

	return getManhattanDistance(w.Ship), nil
}

func getManhattanDistance(s ship.Ship) int {
	return int(math.Abs(float64(s.X)) + math.Abs(float64(s.Y)))
}

func parseInstrs(input []string) ([]instructions.Instruction, error) {
	var instrs []instructions.Instruction

	for _, l := range input {
		instr, err := instructions.Parse(l)
		if err != nil {
			return instrs, err
		}

		instrs = append(instrs, instr)
	}

	return instrs, nil
}
