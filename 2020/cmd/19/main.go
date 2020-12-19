package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	lines, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	res, err := solveA(lines)
	if err != nil {
		fmt.Println("Error solving A:", err)
		return
	}

	fmt.Println("Result A:", res)
}
