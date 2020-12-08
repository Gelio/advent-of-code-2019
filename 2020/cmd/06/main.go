package main

import (
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"fmt"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Error when reading input", err)
	}

	testcases := testcases.SplitTestCaseLines(lines)

	sum := 0

	for _, testcase := range testcases {
		sum += countUniqueLettersInLines(testcase)
	}

	fmt.Println("Result:", sum)
}

func countUniqueLettersInLines(lines []string) int {
	letterFound := make(map[rune]bool)

	for _, line := range lines {
		for _, char := range line {
			letterFound[char] = true
		}
	}

	return len(letterFound)
}
