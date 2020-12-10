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

	nums, err := parseNums(lines)
	if err != nil {
		fmt.Println("Error when parsing lines:", err)
		return
	}

	const preambleLen = 25
	resA, err := solveA(nums, preambleLen)
	if err != nil {
		fmt.Println("Cannot compute A", err)
	} else {
		fmt.Println("Result A:", resA)
	}
}
