package assembler

import (
	"aoc-2020/cmd/20/positionhashmap"
	"aoc-2020/cmd/20/tile"
)

type TileMap [][]*tile.Tile

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

func (tm TileMap) GetCornerTileIDs() []int {
	lastIndex := len(tm) - 1

	return []int{tm[0][0].ID, tm[lastIndex][0].ID, tm[0][lastIndex].ID, tm[lastIndex][lastIndex].ID}
}

func (tm TileMap) GetMapContent() positionhashmap.PositionHashMap {
	tileSideLen := len(tm[0][0].Hashes)
	borderlessTileSideLen := tileSideLen - 2
	sideLen := len(tm) * borderlessTileSideLen

	phm := positionhashmap.New(sideLen)

	for y, row := range tm {
		for x, t := range row {
			// Without borders
			for tileY, tileRow := range t.Hashes[1 : len(t.Hashes)-1] {
				for tileX, c := range tileRow[1 : len(tileRow)-1] {
					if !c {
						continue
					}

					phm[y*borderlessTileSideLen+tileY][x*borderlessTileSideLen+tileX] = true
				}
			}
		}
	}

	return phm
}
