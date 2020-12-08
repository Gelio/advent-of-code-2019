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

func TestCountUniqueLettersInLines(t *testing.T) {
	groups := testcases.SplitTestCaseLines(strings.Split(input, "\n"))
	expectedLetterCounts := []int{3, 3, 3, 1, 1}

	for i := range groups {
		if result := countUniqueLettersInLines(groups[i]); result != expectedLetterCounts[i] {
			t.Errorf("Case %d error: expected %d, got %d", i+1, expectedLetterCounts[i], result)
		}
	}
}
