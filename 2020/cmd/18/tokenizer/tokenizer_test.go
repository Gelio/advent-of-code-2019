package tokenizer_test

import (
	"aoc-2020/cmd/18/tokenizer"
	"aoc-2020/cmd/18/tokenizer/token"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenizer(t *testing.T) {
	line := "2 * 3 + (4 * 5)"
	tokens, err := tokenizer.Tokenize(line)

	require.NoError(t, err, "tokenizing")

	assert.Equal(t, []interface{}{
		token.Num{Value: 2}, token.Times, token.Num{Value: 3}, token.Plus,
		token.LeftParen, token.Num{Value: 4}, token.Times, token.Num{Value: 5}, token.RightParen,
	}, tokens)
}
