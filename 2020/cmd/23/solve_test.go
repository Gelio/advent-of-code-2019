package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimulateCups(t *testing.T) {
	cases := []struct {
		input  string
		moves  int
		result string
	}{
		{
			input:  "389125467",
			moves:  10,
			result: "92658374",
		},
		{
			input:  "389125467",
			moves:  100,
			result: "67384529",
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res, err := solveA(tt.input, tt.moves)

			require.NoError(t, err, "solving A")
			assert.Equal(t, tt.result, res)
		})
	}
}
