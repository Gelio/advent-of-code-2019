package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Error when reading input", err)
		return
	}

	// QUESTION: are such blocks ok for unrelated parts of functions?
	{
		result, err := SolveA(lines)

		if err != nil {
			fmt.Println("Error in A:", err)
			return
		}

		fmt.Println("Result A:", result)
	}

	{
		result, err := SolveB(lines)
		if err != nil {
			fmt.Println("Error in B:", err)
			return
		}

		fmt.Println("Result B:", result)
	}
}
