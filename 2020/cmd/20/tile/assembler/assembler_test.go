package assembler

import (
	"aoc-2020/cmd/20/tile"
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cases = []struct {
	name              string
	getInput          func() ([]string, error)
	expectedCornerIDs []int
}{
	{
		name: "small example",
		getInput: func() ([]string, error) {
			return stdin.ReadLinesFromReader(strings.NewReader(`Tile 2311:
    ..##.#..#.
    ##..#.....
    #...##..#.
    ####.#...#
    ##.##.###.
    ##...#.###
    .#.#.#..##
    ..#....#..
    ###...#.#.
    ..###..###

    Tile 1951:
    #.##...##.
    #.####...#
    .....#..##
    #...######
    .##.#....#
    .###.#####
    ###.##.##.
    .###....#.
    ..#.#..#.#
    #...##.#..

    Tile 1171:
    ####...##.
    #..##.#..#
    ##.#..#.#.
    .###.####.
    ..###.####
    .##....##.
    .#...####.
    #.##.####.
    ####..#...
    .....##...

    Tile 1427:
    ###.##.#..
    .#..#.##..
    .#.##.#..#
    #.#.#.##.#
    ....#...##
    ...##..##.
    ...#.#####
    .#.####.#.
    ..#..###.#
    ..##.#..#.

    Tile 1489:
    ##.#.#....
    ..##...#..
    .##..##...
    ..#...#...
    #####...#.
    #..#.#.#.#
    ...#.#.#..
    ##.#...##.
    ..##.##.##
    ###.##.#..

    Tile 2473:
    #....####.
    #..#.##...
    #.##..#...
    ######.#.#
    .#...#.#.#
    .#########
    .###.#..#.
    ########.#
    ##...##.#.
    ..###.#.#.

    Tile 2971:
    ..#.#....#
    #...###...
    #.#.###...
    ##.##..#..
    .#####..##
    .#..####.#
    #..#.#..#.
    ..####.###
    ..#.#.###.
    ...#.#.#.#

    Tile 2729:
    ...#.#.#.#
    ####.#....
    ..#.#.....
    ....#..#.#
    .##..##.#.
    .#.####...
    ####.#.#..
    ##.####...
    ##..#.##..
    #.##...##.

    Tile 3079:
    #.#.#####.
    .#..######
    ..#.......
    ######....
    ####.#..#.
    .#...#.##.
    #.#####.##
    ..#.###...
    ..#.......
    ..#.###...`))
		},
		expectedCornerIDs: []int{1951, 3079, 2971, 1171},
	},
	{
		name:              "full input",
		getInput:          func() ([]string, error) { return stdin.ReadLinesFromFile("../../input.txt") },
		expectedCornerIDs: []int{3593, 2797, 3517, 3167},
	},
}

func TestAssembler(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			input, err := tt.getInput()

			require.NoError(t, err, "cannot read input")

			rawTiles := testcases.SplitTestCaseLines(input)

			var tiles []tile.Tile
			for _, rawTile := range rawTiles {
				tile, err := tile.Parse(rawTile)
				require.NoError(t, err, "cannot parse tiles")

				tiles = append(tiles, tile)
			}

			tileMap, err := Assemble(tiles)
			require.NoError(t, err, "cannot assemble tiles")

			tileIDs := tileMap.GetCornerTileIDs()

			assert.ElementsMatch(t, tt.expectedCornerIDs, tileIDs)
		})
	}
}

func BenchmarkAssembler(b *testing.B) {
	for _, tt := range cases {
		b.Run(tt.name, func(b *testing.B) {
			input, err := tt.getInput()

			require.NoError(b, err, "cannot read input")

			rawTiles := testcases.SplitTestCaseLines(input)

			var tiles []tile.Tile
			for _, rawTile := range rawTiles {
				tile, err := tile.Parse(rawTile)
				require.NoError(b, err, "cannot parse tiles")

				tiles = append(tiles, tile)
			}

			for i := 0; i < b.N; i++ {
				Assemble(tiles)
			}
		})
	}
}
