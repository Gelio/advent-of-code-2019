package main

import (
	"aoc-2020/internal/testcases"
	"strings"
	"testing"
)

var input = `abc

a
b
c

ab
ac

a
a
a
a

b`
var groups = testcases.SplitTestCaseLines(strings.Split(input, "\n"))

func TestCountUniqueLettersInLines(t *testing.T) {
	expectedLetterCounts := []int{3, 3, 3, 1, 1}

	for i := range groups {
		if result := countUniqueLettersInLines(groups[i]); result != expectedLetterCounts[i] {
			t.Errorf("Case %d error: expected %d, got %d", i+1, expectedLetterCounts[i], result)
		}
	}
}

func TestCountLettersAppearingInAllLines(t *testing.T) {
	expectedResults := []int{3, 0, 1, 1, 1}

	for i := range groups {
		if result := countLettersAppearingInAllLines(groups[i]); result != expectedResults[i] {
			t.Errorf("Case %d error: expected %d, got %d", i+1, expectedResults[i], result)
		}
	}
}
