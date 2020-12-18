package evaluateb_test

import (
	"aoc-2020/cmd/18/evaluate/evaluateb"
	"aoc-2020/cmd/18/tokenizer"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluateTokens(t *testing.T) {
	cases := []struct {
		input          string
		expectedResult int
	}{
		{input: "(1 * 1 + 1) + 1", expectedResult: 3},
		{input: "1 + 2 * 3 + 4 * 5 + 6", expectedResult: 231},
		{input: "1 + (2 * 3) + (4 * (5 + 6))", expectedResult: 51},
		{input: "2 * 3 + (4 * 5)", expectedResult: 46},
		{input: "5 + (8 * 3 + 9 + 3 * 4 * 3)", expectedResult: 1445},
		{input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", expectedResult: 669060},
		{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", expectedResult: 23340},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(tt.input)
			require.NoError(t, err, "tokenizing line")
			res, err := evaluateb.Tokens(tokens)
			require.NoError(t, err, "evaluating tokens")

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
