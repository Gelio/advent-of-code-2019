package assembler

import (
	"aoc-2020/cmd/20/tile"
	"fmt"
	"math"
)

func Assemble(tiles []tile.Tile) (TileMap, error) {
	ta, err := newTileAssembler(len(tiles))
	if err != nil {
		return nil, fmt.Errorf("cannot initialize tile assembler: %w", err)
	}

	for _, t := range tiles {
		ta.variants[t.ID] = t.GetAllVariants()
	}

	for _, variants := range ta.variants {
		for _, v := range variants {
			if ok := ta.tryAssembleWithInitialTile(v); ok {
				return ta.img, nil
			}
		}
	}

	return nil, fmt.Errorf("could not assemble the image")
}

type tileAssembler struct {
	variants    map[int][]tile.Tile
	imgSize     int
	img         TileMap
	usedTileIDs map[int]bool
}

func newTileAssembler(tilesCount int) (tileAssembler, error) {
	var ta tileAssembler
	ta.usedTileIDs = make(map[int]bool)

	ta.imgSize = int(math.Sqrt(float64(tilesCount)))
	if math.Pow(float64(ta.imgSize), float64(2)) != float64(tilesCount) {
		return ta, fmt.Errorf("invalid number of tiles %d, expected a square of an integer", tilesCount)
	}

	ta.img = make(TileMap, ta.imgSize)
	for y := 0; y < ta.imgSize; y++ {
		ta.img[y] = make([]tile.Tile, ta.imgSize)
	}

	ta.variants = make(map[int][]tile.Tile)

	return ta, nil
}

func (ta *tileAssembler) tryAssembleWithInitialTile(t tile.Tile) bool {
	return ta.tryInsertTile(t, 0, 0)
}

func (ta *tileAssembler) tryInsertTile(t tile.Tile, x, y int) bool {
	if x > 0 {
		if !ta.img[y][x-1].MatchesRight(t) {
			return false
		}
	}

	if y > 0 {
		if !ta.img[y-1][x].MatchesBottom(t) {
			return false
		}
	}

	leftmostTile := x == ta.imgSize-1
	bottommostTile := y == ta.imgSize-1

	ta.img[y][x] = t
	ta.usedTileIDs[t.ID] = true

	if finalTile := leftmostTile && bottommostTile; finalTile {
		return true
	}

	if leftmostTile {
		ta.tryVariants(0, y+1)
	} else {
		ta.tryVariants(x+1, y)
	}

	ta.usedTileIDs[t.ID] = false
	// NOTE: no need to reset ta.img, as it contains structs, not pointers

	return false
}

func (ta *tileAssembler) tryVariants(x, y int) bool {
	for tileID, variants := range ta.variants {
		if ta.usedTileIDs[tileID] {
			continue
		}

		for _, variant := range variants {
			if success := ta.tryInsertTile(variant, x, y); success {
				return true
			}
		}
	}

	return false
}
