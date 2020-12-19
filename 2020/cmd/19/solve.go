package main

import (
	"aoc-2020/cmd/19/rule"
	"aoc-2020/cmd/19/statemachine"
	"aoc-2020/internal/testcases"
	"fmt"
)

func solveA(lines []string) (int, error) {
	rules, messages, err := parseLines(lines)
	if err != nil {
		return 0, fmt.Errorf("cannot parse lines: %w", err)
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

func solveB(lines []string) (int, error) {
	rules, messages, err := parseLines(lines)
	if err != nil {
		return 0, fmt.Errorf("cannot parse lines: %w", err)
	}

	// Ultimate hack. Instead of implementing proper recursion, just unwind the expression N times
	const N = 40
	var rule8Options []interface{}
	for i := 1; i <= N; i++ {
		var ruleIDs []int
		for j := 0; j < i; j++ {
			ruleIDs = append(ruleIDs, 42)
		}
		rule8Options = append(rule8Options, rule.Sequence{RuleIDs: ruleIDs})
	}
	rules[8].Content = rule.Alternative{
		Options: rule8Options,
	}

	var rule11Options []interface{}
	for i := 1; i <= N; i++ {
		var ruleIDs []int
		for j := 0; j < i; j++ {
			ruleIDs = append(ruleIDs, 42)
		}
		for j := 0; j < i; j++ {
			ruleIDs = append(ruleIDs, 31)
		}

		rule11Options = append(rule11Options, rule.Sequence{RuleIDs: ruleIDs})
	}
	rules[11].Content = rule.Alternative{
		Options: rule11Options,
	}

	sm, err := statemachine.New(rules)
	if err != nil {
		return 0, fmt.Errorf("error creating state machine: %w", err)
	}

	matchingMessagesCount := 0
	for _, msg := range messages {
		if sm.Matches(msg) {
			matchingMessagesCount++
		}
	}

	return matchingMessagesCount, nil
}

func parseLines(lines []string) ([]rule.Rule, []string, error) {
	sets := testcases.SplitTestCaseLines(lines)
	if len(sets) != 2 {
		return nil, nil, fmt.Errorf("Invalid input. Got %d newline-separated blocks, expected 2", len(sets))
	}

	rawRules := sets[0]
	messages := sets[1]

	rules, err := rule.ParseLines(rawRules)
	if err != nil {
		return nil, nil, fmt.Errorf("Error parsing rules: %w", err)
	}

	return rules, messages, nil
}
