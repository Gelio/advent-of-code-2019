package main

import (
	"aoc-2020/internal/parse"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	cases := []struct {
		input       string
		expectedRes int
	}{
		{
			input: `16
10
15
5
1
11
7
19
6
12
4`,
			expectedRes: 7 * 5,
		},
		{
			input: `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`,
			expectedRes: 22 * 10,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			nums, err := parse.Ints(strings.Split(tt.input, "\n"))

			require.NoError(t, err)

			res := solveA(nums)

			assert.Equal(t, tt.expectedRes, res)
		})
	}
}
