package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	input := strings.Split(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, "\n")

	result, err := SolveA(input)
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}

	expectedResult := 4
	if result != expectedResult {
		t.Errorf("Invalid result. Received %d, expected %d", result, expectedResult)
	}
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		input          string
		expectedResult int
	}{
		{input: `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`,
			expectedResult: 126},
		{input: `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`, expectedResult: 32},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result, err := SolveB(strings.Split(c.input, "\n"))

			if assert.NoError(t, err) {
				assert.Equal(t, c.expectedResult, result)
			}
		})
	}
}
