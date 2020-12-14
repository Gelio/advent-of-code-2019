package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskA(t *testing.T) {
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
			m := newMaskA(tt.mask)

			res := m.Apply(tt.input)

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}

func TestMaskB(t *testing.T) {
	cases := []struct {
		mask            string
		input           int
		expectedResults []int
	}{
		{
			mask:  "000000000000000000000000000000X1001X",
			input: 0b000000000000000000000000000000101010,
			expectedResults: []int{
				0b000000000000000000000000000000011010,
				0b000000000000000000000000000000011011,
				0b000000000000000000000000000000111010,
				0b000000000000000000000000000000111011,
			},
		},
		{
			mask:  "00000000000000000000000000000000X0XX",
			input: 0b000000000000000000000000000000011010,
			expectedResults: []int{
				0b000000000000000000000000000000010000,
				0b000000000000000000000000000000010001,
				0b000000000000000000000000000000010010,
				0b000000000000000000000000000000010011,
				0b000000000000000000000000000000011000,
				0b000000000000000000000000000000011001,
				0b000000000000000000000000000000011010,
				0b000000000000000000000000000000011011,
			},
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			m := newMaskB(tt.mask)

			res := m.Apply(tt.input)

			assert.Equal(t, tt.expectedResults, res)
		})
	}
}
