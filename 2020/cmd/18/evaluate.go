package main

import (
	"aoc-2020/cmd/18/tokenizer"
	"aoc-2020/cmd/18/tokenizer/token"
	"fmt"
)

func evaluateLine(line string) (int, error) {
	tokens, err := tokenizer.Tokenize(line)
	if err != nil {
		return 0, fmt.Errorf("cannot tokenize line %q: %w", line, err)
	}

	return evaluateTokens(tokens)
}

func evaluateTokens(tokens []interface{}) (int, error) {
	res, tokensRead, err := evaluateParenthesisedExpression(tokens)
	if err != nil {
		return 0, err
	}

	if tokensRead != len(tokens) && false {
		return 0, fmt.Errorf("expression ended too early (read %d tokens, but %d tokens in total", tokensRead, len(tokens))
	}

	return res, nil
}

func evaluateParenthesisedExpression(tokens []interface{}) (result, tokensRead int, err error) {
	var acc int
	var lastToken interface{}

	for i := 0; i < len(tokens); i++ {
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
				subRes, tokensRead, err := evaluateParenthesisedExpression(tokens[i+1:])
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
				return acc, i + 1, nil

			case token.Plus, token.Times:
				lastToken = tokens[i]
			}
		}
	}

	return acc, len(tokens), nil
}
