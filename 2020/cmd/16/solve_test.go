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

func TestMatchSpecFromIndex(t *testing.T) {
	specs, err := parseSpecs(strings.Split(`class: 0-1 or 4-19
  row: 0-5 or 8-19
  seat: 0-13 or 16-19`, "\n"))
	require.NoError(t, err)

	nearbyTickets, err := parseTickets(strings.Split(`3,9,18
15,1,5
5,14,9`, "\n"))
	require.NoError(t, err)

	specToFieldMapping := make(map[int]int)
	assert.True(t, matchSpecFromIndex(specs, specToFieldMapping, nearbyTickets, 0), "solution not found")

	assert.Equal(t, map[int]int{
		// class is the 2nd field
		0: 1,
		// row is the 1st field
		1: 0,
		// seat is the 3rd field
		2: 2,
	}, specToFieldMapping)
}
