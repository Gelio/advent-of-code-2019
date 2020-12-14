package main

import "fmt"

func solveA(lines []string) (int, error) {
	memory := make(map[int]int)
	var m maskA

	for i, line := range lines {
		op, err := parseOpLine(line)

		if err != nil {
			return 0, fmt.Errorf("error when parsing line %d: %w", i+1, err)
		}

		switch v := op.(type) {
		case setMask:
			m = newMaskA(v.mask)

		case setMemory:
			memory[v.memIndex] = m.Apply(v.value)
		}
	}

	sum := 0
	for _, x := range memory {
		sum += x
	}

	return sum, nil
}

func solveB(lines []string) (int, error) {
	memory := make(map[int]int)
	var m masksB

	for i, line := range lines {
		op, err := parseOpLine(line)

		if err != nil {
			return 0, fmt.Errorf("error when parsing line %d: %w", i+1, err)
		}

		switch v := op.(type) {
		case setMask:
			m = newMaskB(v.mask)

		case setMemory:
			if m == nil {
				return 0, fmt.Errorf("invalid instruction order, cannot set memory with an unset mask")
			}
			for _, i := range m.Apply(v.memIndex) {
				memory[i] = v.value
			}
		}
	}

	sum := 0
	for _, x := range memory {
		sum += x
	}

	return sum, nil
}
