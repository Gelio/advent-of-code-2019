package main

import (
	"aoc-2020/cmd/18/evaluate/evaluatea"
	"fmt"
)

func solveA(tokenLines [][]interface{}) (int, error) {
	sum := 0

	for _, tokens := range tokenLines {
		res, err := evaluatea.Tokens(tokens)
		if err != nil {
			return 0, fmt.Errorf("cannot evaluate line %q: %w", tokens, err)
		}

		sum += res
	}

	return sum, nil
}
