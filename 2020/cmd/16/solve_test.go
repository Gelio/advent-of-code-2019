package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	specs, err := parseSpecs(strings.Split(`class: 1-3 or 5-7
  row: 6-11 or 33-44
  seat: 13-40 or 45-50`, "\n"))
	require.NoError(t, err)

	nearbyTickets, err := parseTickets(strings.Split(`7,3,47
40,4,50
55,2,20
38,6,12`, "\n"))
	require.NoError(t, err)

	res := solveA(specs, nearbyTickets)
	assert.Equal(t, 71, res)
}
