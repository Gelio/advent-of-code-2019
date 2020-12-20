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
	cases := []struct {
		rules               string
		matchingMessages    []string
		notMatchingMessages []string
	}{
		{
			rules: `0: 1 2
1: "a"
2: 1 3 | 3 1
3: "b"`,
			matchingMessages:    []string{"aab", "aba"},
			notMatchingMessages: []string{"aaaa", "abba", "aaba", "b"},
		},
		{
			rules: `1: 2 3 | 3 2
0: 4 1 5
4: "a"
3: 4 5 | 5 4
2: 4 4 | 5 5
5: "b"`,
			matchingMessages:    []string{"aaaabb", "aaabab", "abbabb", "abbbab", "aabaab", "aabbbb", "abaaab", "ababbb"},
			notMatchingMessages: []string{"bababa", "aaabbb", "aaaabbb", ""},
		},
		{
			rules: `0: 1 2 | 3 3 4 4
1: 3 3 | 4 4
2: 3 4 | 4 3
3: "a"
4: "b"`,
			matchingMessages:    []string{"aaab", "aaba", "bbab", "bbba", "aabb"},
			notMatchingMessages: []string{"aaaa", "bbbb", "bbba", "aabbb"},
		},
		{
			rules: `0: 1 1
1: 2 | 3
2: 4
4: "a"
3: "b"`,
			matchingMessages:    []string{"aa", "ab", "ba", "bb"},
			notMatchingMessages: []string{"a", "b", "aab", "abb"},
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			parsedRules, err := rule.ParseLines(strings.Split(tt.rules, "\n"))
			require.NoError(t, err, "parsing ruless")

			m, err := New(parsedRules)
			require.NoError(t, err, "compiling state machine")

			for i, msg := range tt.matchingMessages {
				t.Run(fmt.Sprintf("matching message %d", i+1), func(t *testing.T) {
					assert.Truef(t, m.Matches(msg), "valid message %q does not match", msg)
				})
			}

			for i, msg := range tt.notMatchingMessages {
				t.Run(fmt.Sprintf("not matching message %d", i+1), func(t *testing.T) {
					assert.Falsef(t, m.Matches(msg), "invalid message %q matches", msg)
				})
			}

		})
	}
}
