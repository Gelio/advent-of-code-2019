package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolveA(t *testing.T) {
	input := strings.Split(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`, "\n")

	var foods []food
	for _, line := range input {
		f, err := parseFood(line)
		require.NoError(t, err, "parsing food")

		foods = append(foods, f)
	}

	res, err := solveA(foods)
	require.NoError(t, err, "solving A")

	assert.Equal(t, 5, res)
}
