package assembler

import "aoc-2020/cmd/20/tile"

type TileMap [][]tile.Tile

func (tm TileMap) GetTileIDs() [][]int {
	var ids [][]int

	for _, row := range tm {
		var idsRow []int

		for _, t := range row {
			idsRow = append(idsRow, t.ID)
		}

		ids = append(ids, idsRow)
	}

	return ids
}
