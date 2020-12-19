package statemachine

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"aoc-2020/cmd/19/rule"
)

func TestStateMachine(t *testing.T) {
	rules := strings.Split(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`, "\n")
	parsedRules, err := rule.ParseLines(rules)
	require.NoError(t, err, "parsing ruless")

	matchingMessages := []string{"ababbb", "abbbab"}
	notMatchingMessages := []string{"bababa", "aaabbb", "aaaabbb"}

	m, err := New(parsedRules)
	require.NoError(t, err, "compiling state machine")

	for i, msg := range matchingMessages {
		t.Run(fmt.Sprintf("matching message %d", i+1), func(t *testing.T) {
			assert.True(t, m.Matches(msg), "valid message does not match")
		})
	}

	for i, msg := range notMatchingMessages {
		t.Run(fmt.Sprintf("not matching message %d", i+1), func(t *testing.T) {
			assert.False(t, m.Matches(msg), "invalid message matches")
		})
	}
}
