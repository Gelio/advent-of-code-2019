package main

import (
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

	res, err := solveA(tokenLines)
	if err != nil {
		fmt.Println("Error when solving A:", err)
		return
	}

	fmt.Println("Result A:", res)

}
