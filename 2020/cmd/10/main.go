package main

import (
	"aoc-2020/internal/parse"
	"aoc-2020/internal/stdin"
	"fmt"
)

func main() {
	input, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Cannot read input:", err)
		return
	}

	nums, err := parse.Ints(input)
	if err != nil {
		fmt.Println("Error when parsing input:", err)
		return
	}

	res := solveA(nums)
	fmt.Println("Result A:", res)

	res = solveB(nums)
	fmt.Println("Result B:", res)
}
