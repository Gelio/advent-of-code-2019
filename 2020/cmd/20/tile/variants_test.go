package tile

import (
	"aoc-2020/internal/stdin"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllVariants(t *testing.T) {
	cases := []struct {
		input         string
		variantsCount int
	}{
		{input: "Tile 1:\n123\n456\n789", variantsCount: 8},
		{input: "Tile 2:\n111\n111\n111", variantsCount: 1},
		{input: "Tile 3:\n111\n222\n111", variantsCount: 2},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			tile, err := Parse(strings.Split(tt.input, "\n"))

			require.NoError(t, err, "parsing tile")

			variants := tile.GetAllVariants()

			assert.Len(t, variants, tt.variantsCount, "invalid number of variants")
		})
	}
}

func TestMatchingExistingVariants(t *testing.T) {
	cases := []struct {
		ID          int
		input       string
		tileVariant string
	}{
		{
			ID: 1951,
			input: `Tile 1951:
      #.##...##.
      #.####...#
      .....#..##
      #...######
      .##.#....#
      .###.#####
      ###.##.##.
      .###....#.
      ..#.#..#.#
      #...##.#..`,
			tileVariant: `Tile 1951:
      #...##.#..
      ..#.#..#.#
      .###....#.
      ###.##.##.
      .###.#####
      .##.#....#
      #...######
      .....#..##
      #.####...#
      #.##...##.`,
		},
		{
			ID: 2311,
			input: `Tile 2311:
      ..##.#..#.
      ##..#.....
      #...##..#.
      ####.#...#
      ##.##.###.
      ##...#.###
      .#.#.#..##
      ..#....#..
      ###...#.#.
      ..###..###`,
			tileVariant: `Tile 2311:
      ..###..###
      ###...#.#.
      ..#....#..
      .#.#.#..##
      ##...#.###
      ##.##.###.
      ####.#...#
      #...##..#.
      ##..#.....
      ..##.#..#.`,
		},
		{
			ID: 3079,
			input: `Tile 3079:
      #.#.#####.
      .#..######
      ..#.......
      ######....
      ####.#..#.
      .#...#.##.
      #.#####.##
      ..#.###...
      ..#.......
      ..#.###...`,
			tileVariant: `Tile 3079:
      #.#.#####.
      .#..######
      ..#.......
      ######....
      ####.#..#.
      .#...#.##.
      #.#####.##
      ..#.###...
      ..#.......
      ..#.###...`,
		},
		{
			ID: 2729,
			input: `Tile 2729:
      ...#.#.#.#
      ####.#....
      ..#.#.....
      ....#..#.#
      .##..##.#.
      .#.####...
      ####.#.#..
      ##.####...
      ##..#.##..
      #.##...##.`,
			tileVariant: `Tile 2729:
      #.##...##.
      ##..#.##..
      ##.####...
      ####.#.#..
      .#.####...
      .##..##.#.
      ....#..#.#
      ..#.#.....
      ####.#....
      ...#.#.#.#`,
		},
		{
			ID: 1427,
			input: `Tile 1427:
      ###.##.#..
      .#..#.##..
      .#.##.#..#
      #.#.#.##.#
      ....#...##
      ...##..##.
      ...#.#####
      .#.####.#.
      ..#..###.#
      ..##.#..#.`,
			tileVariant: `Tile 1427:
      ..##.#..#.
      ..#..###.#
      .#.####.#.
      ...#.#####
      ...##..##.
      ....#...##
      #.#.#.##.#
      .#.##.#..#
      .#..#.##..
      ###.##.#..`,
		},
		{
			ID: 2473,
			input: `Tile 2473:
      #....####.
      #..#.##...
      #.##..#...
      ######.#.#
      .#...#.#.#
      .#########
      .###.#..#.
      ########.#
      ##...##.#.
      ..###.#.#.`,
			tileVariant: `Tile 2473:
      ..#.###...
      ##.##....#
      ..#.###..#
      ###.#..###
      .######.##
      #.#.#.#...
      #.###.###.
      #.###.##..
      .######...
      .##...####`,
		},
		{
			ID: 2971,
			input: `Tile 2971:
      ..#.#....#
      #...###...
      #.#.###...
      ##.##..#..
      .#####..##
      .#..####.#
      #..#.#..#.
      ..####.###
      ..#.#.###.
      ...#.#.#.#`,
			tileVariant: `Tile 2971:
      ...#.#.#.#
      ..#.#.###.
      ..####.###
      #..#.#..#.
      .#..####.#
      .#####..##
      ##.##..#..
      #.#.###...
      #...###...
      ..#.#....#`,
		},
		{
			ID: 1489,
			input: `Tile 1489:
      ##.#.#....
      ..##...#..
      .##..##...
      ..#...#...
      #####...#.
      #..#.#.#.#
      ...#.#.#..
      ##.#...##.
      ..##.##.##
      ###.##.#..`,
			tileVariant: `Tile 1489:
      ###.##.#..
      ..##.##.##
      ##.#...##.
      ...#.#.#..
      #..#.#.#.#
      #####...#.
      ..#...#...
      .##..##...
      ..##...#..
      ##.#.#....`,
		},
		{
			ID: 1171,
			input: `Tile 1171:
      ####...##.
      #..##.#..#
      ##.#..#.#.
      .###.####.
      ..###.####
      .##....##.
      .#...####.
      #.##.####.
      ####..#...
      .....##...`,
			tileVariant: `Tile 1171:
      .##...####
      #..#.##..#
      .#.#..#.##
      .####.###.
      ####.###..
      .##....##.
      .####...#.
      .####.##.#
      ...#..####
      ...##.....`,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("tile %d", tt.ID), func(t *testing.T) {
			lines, err := stdin.ReadLinesFromReader(strings.NewReader(tt.input))
			require.NoError(t, err, "reading input")

			parsedTile, err := Parse(lines)
			require.NoError(t, err, "parsing tile")

			lines, err = stdin.ReadLinesFromReader(strings.NewReader(tt.tileVariant))
			require.NoError(t, err, "reading input")

			finalTile, err := Parse(lines)
			require.NoError(t, err, "parsing final tile")

			for _, variant := range parsedTile.GetAllVariants() {
				if variant.Borders == finalTile.Borders {
					t.Logf("tile %d matches: %v %v", parsedTile.ID, variant.Rotation, variant.Flip)
					return
				}
			}

			assert.Fail(t, "did not match tile variant")
		})
	}
}
