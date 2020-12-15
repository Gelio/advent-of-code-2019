package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveA(t *testing.T) {
	// Only cases from A are tested, since part B is just related to performance optimizations,
	// not correctness.
	resultIndex := 2020
	cases := []struct {
		input          []int
		expectedResult int
	}{
		{
			input:          []int{0, 3, 6},
			expectedResult: 436,
		},
		{
			input:          []int{1, 3, 2},
			expectedResult: 1,
		},
		{
			input:          []int{2, 1, 3},
			expectedResult: 10,
		},
		{
			input:          []int{1, 2, 3},
			expectedResult: 27,
		},
		{
			input:          []int{2, 3, 1},
			expectedResult: 78,
		},
		{
			input:          []int{3, 2, 1},
			expectedResult: 438,
		},
		{
			input:          []int{3, 1, 2},
			expectedResult: 1836,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := solve(tt.input, resultIndex)

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
