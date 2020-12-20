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
	assert.Equal(t, tile.Borders.Top, "##.#..##.#", "invalid top border")
	assert.Equal(t, tile.Borders.Bottom, "...##..#..", "invalid bottom border")
	assert.Equal(t, tile.Borders.Left, "#....#.##.", "invalid left border")
	assert.Equal(t, tile.Borders.Right, "#.#..####.", "invalid right border")
}
