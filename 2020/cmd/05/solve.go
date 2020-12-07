package main

import (
	"fmt"
)

func solve(input []string) int {
	return 0
}

const (
	lower = iota
	upper = iota
)

var characters = map[rune]int{
	'F': lower,
	'B': upper,
	'R': upper,
	'L': lower,
}

func mapCharacters(input string) (ops []int) {
	ops = make([]int, len(input))
	for i, b := range input {
		op, ok := characters[b]
		if !ok {
			panic(fmt.Errorf("Invalid operation: %v", b))
		}

		ops[i] = op
	}

	return
}

func getPosition(ops []int) int {
	min := 0
	max := 1 << len(ops)

	for _, op := range ops {
		middle := min + (max-min-1)/2 + 1

		if op == upper {
			min = middle
		} else if op == lower {
			max = middle
		} else {
			panic(fmt.Errorf("Invalid op: %d", op))
		}
	}

	return min
}

type passengerSeat struct {
	row, column int
}

func getPassengerSeat(input string) (s passengerSeat, err error) {
	const (
		rowPosLen    = 7
		columnPosLen = 3
	)
	if len(input) != rowPosLen+columnPosLen {
		err = fmt.Errorf("Invalid input has length of %d, expected %d", len(input), rowPosLen+columnPosLen)
		return
	}

	rowPos := input[0:rowPosLen]
	s.row = getPosition(mapCharacters(rowPos))
	colPos := input[rowPosLen:]
	s.column = getPosition(mapCharacters(colPos))

	return
}

func getSeatID(s *passengerSeat) int {
	return s.row*8 + s.column
}
