package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	input, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading input:", err)
		return
	}

	resA, err := solveA(input)
	if err != nil {
		fmt.Println("Error while solving A:", err)
		return
	}

	fmt.Println("Result A:", resA)
}
