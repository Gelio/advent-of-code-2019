package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	input, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var foods []food
	for _, line := range input {
		f, err := parseFood(line)
		if err != nil {
			fmt.Println("Error parsing food:", err)
			return
		}

		foods = append(foods, f)
	}

	res, err := solveA(foods)
	if err != nil {
		fmt.Println("Error solving A:", err)
		return
	}
	fmt.Println("Result A:", res)
}
