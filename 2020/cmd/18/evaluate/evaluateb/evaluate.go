package evaluateb

import (
	"aoc-2020/cmd/18/tokenizer/token"
	"fmt"
)

func Tokens(tokens []interface{}) (int, error) {
	res, tokensRead, err := evaluateUntil(tokens, nil)
	if err != nil {
		return 0, err
	}

	if tokensRead != len(tokens) && false {
		return 0, fmt.Errorf("expression ended too early (read %d tokens, but %d tokens in total", tokensRead, len(tokens))
	}

	return res, nil
}

func evaluateUntil(tokens []interface{}, endToken interface{}) (result, tokensRead int, err error) {
	var acc int
	var lastToken interface{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i] == endToken {
			return acc, i, nil
		}

		switch t := tokens[i].(type) {
		case token.Num:
			if lastToken == token.Plus || lastToken == nil {
				acc += t.Value
			} else if lastToken == token.Times {
				acc *= t.Value
			} else {
				return 0, i, fmt.Errorf("unexpected number %d after %#v at token index %d", t.Value, lastToken, i)
			}
			lastToken = tokens[i]

		case rune:
			switch t {
			case token.LeftParen:
				subRes, tokensRead, err := evaluateUntil(tokens[i+1:], nil)
				if err != nil {
					return 0, i, fmt.Errorf("cannot evaluate parenthesised expression from token index %d: %w", i, err)
				}
				if lastToken == token.Plus || lastToken == nil {
					acc += subRes
				} else if lastToken == token.Times {
					acc *= subRes
				}
				lastToken = token.Num{Value: subRes}
				i += tokensRead

			case token.RightParen:
				if endToken == nil || endToken == token.RightParen {
					return acc, i + 1, nil
				}

				return acc, i, nil

			case token.Plus:
				lastToken = tokens[i]

			case token.Times:
				subRes, tokensRead, err := evaluateUntil(tokens[i+1:], token.Times)
				if err != nil {
					return 0, i, fmt.Errorf("cannot evaluate parenthesised expression from token index %d: %w", i, err)
				}
				acc *= subRes
				lastToken = token.Times
				i += tokensRead
			}
		}
	}

	return acc, len(tokens), nil
}
