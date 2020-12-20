package tile

import (
	"aoc-2020/internal/stdin"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	input, err := stdin.ReadLinesFromReader(
		strings.NewReader(`Tile 1367:
      ##.#..##.#
      ..#.......
      .#.......#
      ...#..##..
      .##....##.
      #..#.##..#
      ..#...#..#
      #.#......#
      #..#....##
      ...##..#..`),
	)
	require.NoError(t, err, "parsing input")

	tile, err := Parse(input)

	assert.Equal(t, tile.ID, 1367, "invalid tile ID")
	assert.True(t, tile.Hashes[0][0])
	assert.True(t, tile.Hashes[0][1])
	assert.False(t, tile.Hashes[0][2])
	assert.False(t, tile.Hashes[1][0])
}
