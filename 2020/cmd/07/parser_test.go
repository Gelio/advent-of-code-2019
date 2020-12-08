package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	cases := []struct {
		line     string
		expected *rule
	}{
		{line: "light red bags contain 1 bright white bag, 2 muted yellow bags.", expected: &rule{
			bagColor: "light red",
			contents: []*bagWithQuantity{
				{color: "bright white", quantity: 1},
				{color: "muted yellow", quantity: 2},
			},
		}},
		{line: "light red bags contain 1 bright white bag, 2 muted yellow bags, 1 shiny gold bag.", expected: &rule{
			bagColor: "light red",
			contents: []*bagWithQuantity{
				{color: "bright white", quantity: 1},
				{color: "muted yellow", quantity: 2},
				{color: "shiny gold", quantity: 1},
			},
		}},
		{line: "bright white bags contain 1 shiny gold bag.", expected: &rule{
			bagColor: "bright white",
			contents: []*bagWithQuantity{
				{color: "shiny gold", quantity: 1},
			},
		}},
		{line: "faded blue bags contain no other bags.", expected: &rule{
			bagColor: "faded blue",
			contents: []*bagWithQuantity{},
		}},
	}

	for caseIndex, c := range cases {
		result, err := parseLine(c.line)
		if err != nil {
			t.Errorf("Case %d: error %v", caseIndex+1, err)
			continue
		}

		if result.bagColor != c.expected.bagColor {
			t.Errorf("Case %d: invalid bag color, expected %#v, got %#v", caseIndex+1, c.expected.bagColor, result.bagColor)
		}

		if len(result.contents) != len(c.expected.contents) {
			t.Errorf("Case %d: invalid number of contents, expected %v, got %v", caseIndex+1, len(c.expected.contents), len(result.contents))
			continue
		}

		for i, bag := range result.contents {
			if expectedColor := c.expected.contents[i].color; bag.color != expectedColor {
				t.Errorf("Case %d, bag %d: invalid color, expected %#v, got %#v", caseIndex+1, i, bag.color, expectedColor)
			}

			if expectedQuantity := c.expected.contents[i].quantity; bag.quantity != expectedQuantity {
				t.Errorf("Case %d, bag %d: invalid quantity, expected %v, got %v", caseIndex+1, i, bag.quantity, expectedQuantity)
			}
		}
	}
}
