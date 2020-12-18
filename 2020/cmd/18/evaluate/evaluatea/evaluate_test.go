package evaluatea_test

import (
	"aoc-2020/cmd/18/evaluate/evaluatea"
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
		{input: "1 + 2 * 3 + 4 * 5 + 6", expectedResult: 71},
		{input: "1 + (2 * 3) + (4 * (5 + 6))", expectedResult: 51},
		{input: "2 * 3 + (4 * 5)", expectedResult: 26},
		{input: "5 + (8 * 3 + 9 + 3 * 4 * 3)", expectedResult: 437},
		{input: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", expectedResult: 12240},
		{input: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", expectedResult: 13632},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(tt.input)
			require.NoError(t, err, "tokenizing line")
			res, err := evaluatea.Tokens(tokens)
			require.NoError(t, err, "evaluating tokens")

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
