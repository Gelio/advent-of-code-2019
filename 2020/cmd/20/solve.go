package main

import (
	"aoc-2020/cmd/20/monsterfinder"
	"aoc-2020/cmd/20/positionhashmap"
	"aoc-2020/cmd/20/tile/assembler"
)

func solveA(tileMap assembler.TileMap) int {
	cornerIDs := tileMap.GetCornerTileIDs()
	res := 1
	for _, id := range cornerIDs {
		res *= id
	}

	return res
}

func solveB(phm positionhashmap.PositionHashMap) int {
	monstersCount := monsterfinder.GetMonstersCount(phm)

	return phm.CountValues() - monstersCount*15
}
