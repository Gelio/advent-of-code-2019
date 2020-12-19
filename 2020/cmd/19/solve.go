package main

import (
	"aoc-2020/cmd/19/rule"
	"aoc-2020/cmd/19/statemachine"
	"aoc-2020/internal/testcases"
	"fmt"
)

func solveA(lines []string) (int, error) {
	sets := testcases.SplitTestCaseLines(lines)
	if len(sets) != 2 {
		return 0, fmt.Errorf("Invalid input. Got %d newline-separated blocks, expected 2", len(sets))
	}

	rawRules := sets[0]
	messages := sets[1]

	rules, err := rule.ParseLines(rawRules)
	if err != nil {
		return 0, fmt.Errorf("Error parsing rules: %w", err)
	}

	sm, err := statemachine.New(rules)
	if err != nil {
		return 0, fmt.Errorf("Error creating state machine: %w", err)
	}

	matchingMessagesCount := 0
	for _, msg := range messages {
		if sm.Matches(msg) {
			matchingMessagesCount++
		}
	}

	return matchingMessagesCount, nil
}
