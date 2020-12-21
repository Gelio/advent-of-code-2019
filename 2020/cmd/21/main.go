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

	fmt.Println("Result A:", solveA(foods))
	fmt.Println("Result B:", solveB(foods))
}
