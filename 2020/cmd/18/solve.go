package main

import (
	"fmt"
)

func solve(tokenLines [][]interface{}, evaluate func(tokens []interface{}) (int, error)) (int, error) {
	sum := 0

	for _, tokens := range tokenLines {
		res, err := evaluate(tokens)
		if err != nil {
			return 0, fmt.Errorf("cannot evaluate line %q: %w", tokens, err)
		}

		sum += res
	}

	return sum, nil
}
