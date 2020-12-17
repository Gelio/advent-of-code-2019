package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Error when reading input:", err)
		return
	}

	res, err := solve(lines, getNeighboringPositions3D)
	if err != nil {
		fmt.Println("Error when solving A:", err)
		return
	}

	fmt.Println("Result A:", res)

	res, err = solve(lines, getNeighboringPositions4D)
	if err != nil {
		fmt.Println("Error when solving B:", err)
		return
	}

	fmt.Println("Result B:", res)
}
