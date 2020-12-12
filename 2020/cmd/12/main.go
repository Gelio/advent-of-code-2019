package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	lines, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error when reading input:", err)
		return
	}

	res, err := solveA(lines)
	if err != nil {
		fmt.Println("Error when solving A:", err)
		return
	}

	fmt.Println("Result A:", res)

	res, err = solveB(lines)
	if err != nil {
		fmt.Println("Error when solving B:", err)
		return
	}

	fmt.Println("Result B:", res)
}
