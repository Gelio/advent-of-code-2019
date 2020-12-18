package main

import (
	"aoc-2020/cmd/18/evaluate/evaluatea"
	"aoc-2020/cmd/18/evaluate/evaluateb"
	"aoc-2020/cmd/18/tokenizer"
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Error when reading input:", err)
		return
	}

	var tokenLines [][]interface{}

	for _, line := range lines {
		tokens, err := tokenizer.Tokenize(line)
		if err != nil {
			fmt.Println("Error when tokenizing line", line, err)
			return
		}

		tokenLines = append(tokenLines, tokens)
	}

	res, err := solve(tokenLines, evaluatea.Tokens)
	if err != nil {
		fmt.Println("Error when solving A:", err)
		return
	}

	fmt.Println("Result A:", res)

	res, err = solve(tokenLines, evaluateb.Tokens)
	if err != nil {
		fmt.Println("Error when solving B:", err)
		return
	}

	fmt.Println("Result B:", res)
}
