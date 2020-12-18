package tokenizer

import (
	"aoc-2020/cmd/18/tokenizer/token"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

func Tokenize(line string) ([]interface{}, error) {
	var tokens []interface{}
	var s scanner.Scanner
	s.Init(strings.NewReader(line))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case scanner.Int:
			t := s.TokenText()
			num, err := strconv.Atoi(t)
			if err != nil {
				return nil, fmt.Errorf("cannot parse int %q at position %d: %w", t, s.Position.Column, err)
			}

			tokens = append(tokens, token.Num{Value: num})

		case '(':
			tokens = append(tokens, token.LeftParen)
		case ')':
			tokens = append(tokens, token.RightParen)
		case '+':
			tokens = append(tokens, token.Plus)
		case '*':
			tokens = append(tokens, token.Times)

		default:
			return nil, fmt.Errorf("unknown token %q at position %d", tok, s.Position.Column)
		}
	}

	return tokens, nil
}
