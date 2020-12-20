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

	var variants []tile.Tile
	for _, t := range tiles {
		for _, v := range t.GetAllVariants() {
			ta.addTile(v)
			variants = append(variants, v)
		}
	}

	for _, v := range variants {
		if ok := ta.tryAssembleWithInitialTile(v); ok {
			return ta.img, nil
		}
	}

	return nil, fmt.Errorf("could not assemble the image")
}

type tileAssembler struct {
	borders     map[string][]tile.Tile
	imgSize     int
	img         TileMap
	usedTileIDs map[int]bool
}

func newTileAssembler(tilesCount int) (tileAssembler, error) {
	var ta tileAssembler
	ta.borders = make(map[string][]tile.Tile)
	ta.usedTileIDs = make(map[int]bool)

	ta.imgSize = int(math.Sqrt(float64(tilesCount)))
	if math.Pow(float64(ta.imgSize), float64(2)) != float64(tilesCount) {
		return ta, fmt.Errorf("invalid number of tiles %d, expected a square of an integer", tilesCount)
	}

	ta.img = make(TileMap, ta.imgSize)
	for y := 0; y < ta.imgSize; y++ {
		ta.img[y] = make([]tile.Tile, ta.imgSize)
	}

	return ta, nil
}

func (ta *tileAssembler) addTile(t tile.Tile) {
	ta.addBorder(t.Borders.Top, t)
	ta.addBorder(t.Borders.Bottom, t)
	ta.addBorder(t.Borders.Left, t)
	ta.addBorder(t.Borders.Right, t)
}

func (ta *tileAssembler) addBorder(border string, t tile.Tile) {
	ta.borders[border] = append(ta.borders[border], t)
}

func (ta *tileAssembler) tryAssembleWithInitialTile(t tile.Tile) bool {
	return ta.tryInsertTile(t, 0, 0)
}

func (ta *tileAssembler) tryInsertTile(t tile.Tile, x, y int) bool {
	// if x > 0 {
	// 	if matchesTileLeft := ta.img[y][x-1].Borders.Right == t.Borders.Left; !matchesTileLeft {
	// 		return false
	// 	}
	// }

	// if y > 0 {
	// 	if matchesTileUp := ta.img[y-1][x].Borders.Bottom == t.Borders.Top; !matchesTileUp {
	// 		return false
	// 	}
	// }

	leftmostTile := x == ta.imgSize-1
	if !leftmostTile && !ta.hasPossiblyBorderingTiles(t.Borders.Right, t.ID) {
		return false
	}

	bottommostTile := y == ta.imgSize-1
	if !bottommostTile && !ta.hasPossiblyBorderingTiles(t.Borders.Bottom, t.ID) {
		return false
	}

	ta.img[y][x] = t
	ta.usedTileIDs[t.ID] = true

	if finalTile := leftmostTile && bottommostTile; finalTile {
		return true
	}

	if leftmostTile {
		firstTileInRow := ta.img[y][0]
		for _, tileCandidate := range ta.getPossiblyBorderingTiles(firstTileInRow.Borders.Bottom, firstTileInRow.ID) {
			if success := ta.tryInsertTile(tileCandidate, 0, y+1); success {
				return true
			}
		}
	} else {
		for _, tileCandidate := range ta.getPossiblyBorderingTiles(t.Borders.Right, t.ID) {
			if success := ta.tryInsertTile(tileCandidate, x+1, y); success {
				return true
			}
		}
	}

	ta.usedTileIDs[t.ID] = false
	// NOTE: no need to reset ta.img, as it contains structs, not pointers

	return false
}

func (ta *tileAssembler) hasPossiblyBorderingTiles(border string, tileID int) bool {
	for _, t := range ta.borders[border] {
		if !ta.usedTileIDs[t.ID] && t.ID != tileID {
			return true
		}
	}

	return false
}

func (ta *tileAssembler) getPossiblyBorderingTiles(border string, tileID int) []tile.Tile {
	var borderingTiles []tile.Tile
	for _, t := range ta.borders[border] {
		if !ta.usedTileIDs[t.ID] && t.ID != tileID {
			borderingTiles = append(borderingTiles, t)
		}
	}

	return borderingTiles
}
