package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	cases := []struct {
		mask                  string
		input, expectedResult int
	}{
		{mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", input: 11, expectedResult: 73},
		{mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", input: 101, expectedResult: 101},
		{mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", input: 0, expectedResult: 64},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			m := newMask(tt.mask)

			res := m.Apply(tt.input)

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
