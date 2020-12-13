package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBusIDs(t *testing.T) {
	busIDs, err := parseBusIDs("7,13,x,x,59,x,31,19")

	require.NoError(t, err)

	assert.Equal(t, busIDs, []int{7, 13, cross, cross, 59, cross, 31, 19})
}

func TestSolveA(t *testing.T) {
	timestamp := 939
	busIDs := []int{7, 13, 59, 31, 19}

	result := solveA(timestamp, busIDs)

	assert.Equal(t, 295, result)
}

func TestSolveB(t *testing.T) {
	cases := []struct {
		busIDs         string
		startTimestamp int
		expectedResult int
	}{
		{
			busIDs:         "7,13,x,x,59,x,31,19",
			startTimestamp: 0,
			expectedResult: 1068781,
		},
		{
			busIDs:         "17,x,13,19",
			startTimestamp: 0,
			expectedResult: 3417,
		},
		{
			busIDs:         "67,7,59,61",
			startTimestamp: 0,
			expectedResult: 754018,
		},
		{
			busIDs:         "67,x,7,59,61",
			startTimestamp: 0,
			expectedResult: 779210,
		},
		{
			busIDs:         "67,7,x,59,61",
			startTimestamp: 0,
			expectedResult: 1261476,
		},
		{
			busIDs:         "1789,37,47,1889",
			startTimestamp: 0,
			expectedResult: 1202161486,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			busIDs, err := parseBusIDs(tt.busIDs)

			require.NoError(t, err, "Invalid bus IDs in the test case")

			result := solveB(tt.startTimestamp, busIDs)
			assert.Equal(t, tt.expectedResult, result)
		})
	}

}
