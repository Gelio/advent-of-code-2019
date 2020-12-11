package main

import (
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	input, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Cannot read input:", err)
		return
	}

	res := solveA(input)
	fmt.Println("Result A:", res)

	res = solveB(input)
	fmt.Println("Result B:", res)
}
