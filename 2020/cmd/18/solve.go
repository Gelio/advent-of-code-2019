package main

import (
	"fmt"
)

func solveA(lines []string) (int, error) {
	sum := 0

	for _, line := range lines {
		res, err := evaluateLine(line)
		if err != nil {
			return 0, fmt.Errorf("cannot evaluate line %q: %w", line, err)
		}

		sum += res
	}

	return sum, nil
}
