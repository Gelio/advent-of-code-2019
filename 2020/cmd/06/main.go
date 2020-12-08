package main

import (
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"fmt"

	"github.com/golang-collections/collections/set"
)

func main() {
	lines, err := stdin.ReadAllLines()
	if err != nil {
		fmt.Println("Error when reading input", err)
	}

	testcases := testcases.SplitTestCaseLines(lines)

	sumA := 0
	sumB := 0

	for _, testcase := range testcases {
		sumA += countUniqueLettersInLines(testcase)
		sumB += countLettersAppearingInAllLines(testcase)
	}

	fmt.Println("Result A:", sumA)
	fmt.Println("Result B:", sumB)
}

func countUniqueLettersInLines(lines []string) int {
	var lettersSet *set.Set

	for _, line := range lines {
		currentLineLettersSet := getLettersSet(line)
		if lettersSet == nil {
			lettersSet = currentLineLettersSet
		} else {
			lettersSet = lettersSet.Union(currentLineLettersSet)
		}
	}

	return lettersSet.Len()
}

func getLettersSet(line string) *set.Set {
	s := set.New()

	for _, char := range line {
		s.Insert(char)
	}

	return s
}

func countLettersAppearingInAllLines(lines []string) int {
	var lettersSet *set.Set

	for _, line := range lines {
		currentLineLettersSet := getLettersSet(line)
		if lettersSet == nil {
			lettersSet = currentLineLettersSet
		} else {
			lettersSet = lettersSet.Intersection(currentLineLettersSet)
		}
	}

	return lettersSet.Len()
}
