package main

import (
	"aoc-2020/cmd/20/tile"
	"aoc-2020/cmd/20/tile/assembler"
	"aoc-2020/internal/stdin"
	"aoc-2020/internal/testcases"
	"fmt"
)

func main() {
	input, err := stdin.ReadLinesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var tiles []tile.Tile
	for _, c := range testcases.SplitTestCaseLines(input) {
		t, err := tile.Parse(c)
		if err != nil {
			fmt.Println("Error parsing tile:", err)
			return
		}

		tiles = append(tiles, t)
	}

	tileMap, err := assembler.Assemble(tiles)
	if err != nil {
		fmt.Println("Error while assembling image:", err)
		return
	}

	fmt.Println("Found", len(tileMap))
}
