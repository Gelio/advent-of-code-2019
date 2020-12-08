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
	letterFound := make(map[rune]bool)

	for _, line := range lines {
		for _, char := range line {
			letterFound[char] = true
		}
	}

	return len(letterFound)
}

func getLettersSet(line string) (s map[rune]bool) {
	s = make(map[rune]bool)

	for _, char := range line {
		s[char] = true
	}

	return
}

func intersectSets(sets []map[rune]bool) (r map[rune]bool) {
	r = make(map[rune]bool)

	if len(sets) == 0 {
		return
	}

	for char := range sets[0] {
		r[char] = true
	}

	for _, set := range sets[1:] {
		for char := range r {
			if ok := set[char]; !ok {
				delete(r, char)
			}
		}
	}

	return
}

func countLettersAppearingInAllLines(lines []string) int {
	var letterSets []map[rune]bool

	for _, line := range lines {
		letterSets = append(letterSets, getLettersSet(line))
	}

	return len(intersectSets(letterSets))
}
