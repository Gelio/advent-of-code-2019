package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOpLine(t *testing.T) {
	cases := []struct {
		line           string
		expectedResult interface{}
	}{
		{
			line:           "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			expectedResult: setMask{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		},
		{
			line:           "mem[8] = 11",
			expectedResult: setMemory{8, 11},
		},
		{
			line:           "mem[7] = 101",
			expectedResult: setMemory{7, 101},
		},
		{
			line:           "mem[8] = 0",
			expectedResult: setMemory{8, 0},
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d (%s)", i+1, tt.line), func(t *testing.T) {
			result, err := parseOpLine(tt.line)

			require.NoError(t, err)

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
