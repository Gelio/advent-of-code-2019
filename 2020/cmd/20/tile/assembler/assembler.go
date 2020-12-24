package assembler

import (
	"aoc-2020/cmd/20/tile"
	"fmt"
	"math"
	"runtime"
)

func Assemble(tiles []tile.Tile) (TileMap, error) {
	ta, err := newTileAssembler(tiles)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize tile assembler: %w", err)
	}

	// Multiple workers are initialized. Each tries to assemble the image starting with a given tile
	// ID. If it succeeds, it signals that via a channel. If not, it signals it is ready for more work
	workersCount := runtime.GOMAXPROCS(0)
	workerAssemblers := []*tileAssembler{&ta}
	workerReadyChan := make(chan int, workersCount)
	workerReadyChan <- 0
	assemblySuccessChan := make(chan int)
	for i := 1; i < workersCount; i++ {
		taClone := ta.clone()
		taClone.ID = i
		workerAssemblers = append(workerAssemblers, &taClone)
		workerReadyChan <- i
	}

	for tileID := range ta.variants {
		select {
		case workerID := <-workerReadyChan:
			ta := workerAssemblers[workerID]
			go ta.tryAssemble(tileID, workerReadyChan, assemblySuccessChan)

		case workerID := <-assemblySuccessChan:
			ta := workerAssemblers[workerID]
			// TODO: cancel other goroutines?
			return ta.img, nil
		}
	}

	return nil, fmt.Errorf("could not assemble the image")
}

type tileAssembler struct {
	ID          int
	variants    map[int][]*tile.Tile
	imgSize     int
	img         TileMap
	usedTileIDs map[int]bool
	// Cache tiles with corresponding variants to consider only tiles with matching borders when trying
	// to match a new tile
	leftBorderVariants map[string][]*tile.Tile
	topBorderVariants  map[string][]*tile.Tile
}

type assemblingResult struct {
	assemblerID int
	success     bool
}

func newTileAssembler(tiles []tile.Tile) (tileAssembler, error) {
	var ta tileAssembler
	tilesCount := len(tiles)
	ta.usedTileIDs = make(map[int]bool)
	ta.leftBorderVariants = make(map[string][]*tile.Tile)
	ta.topBorderVariants = make(map[string][]*tile.Tile)

	ta.imgSize = int(math.Sqrt(float64(tilesCount)))
	if math.Pow(float64(ta.imgSize), float64(2)) != float64(tilesCount) {
		return ta, fmt.Errorf("invalid number of tiles %d, expected a square of an integer", tilesCount)
	}

	ta.img = make(TileMap, ta.imgSize)
	for y := 0; y < ta.imgSize; y++ {
		ta.img[y] = make([]*tile.Tile, ta.imgSize)
	}

	ta.variants = make(map[int][]*tile.Tile)
	for _, t := range tiles {
		variants := t.GetAllVariants()
		ta.variants[t.ID] = variants

		for _, variant := range variants {
			ta.leftBorderVariants[variant.Left] = append(ta.leftBorderVariants[variant.Left], variant)
			ta.topBorderVariants[variant.Top] = append(ta.topBorderVariants[variant.Top], variant)
		}
	}

	return ta, nil
}

func (ta tileAssembler) clone() tileAssembler {
	// Clone only the maps that will be used for reading and writing in concurrent goroutines
	ta.usedTileIDs = make(map[int]bool)
	ta.img = make(TileMap, ta.imgSize)
	for y := 0; y < ta.imgSize; y++ {
		ta.img[y] = make([]*tile.Tile, ta.imgSize)
	}

	return ta
}

func (ta *tileAssembler) tryAssemble(startingTileID int, readyChan chan<- int, successChan chan<- int) {
	for _, t := range ta.variants[startingTileID] {
		if ok := ta.insertTile(t, 0, 0); ok {
			successChan <- ta.ID
			break
		}
	}

	readyChan <- ta.ID
}

func (ta *tileAssembler) insertTile(t *tile.Tile, x, y int) bool {
	rightmostTile := x == ta.imgSize-1
	bottommostTile := y == ta.imgSize-1

	ta.img[y][x] = t
	ta.usedTileIDs[t.ID] = true

	if finalTile := rightmostTile && bottommostTile; finalTile {
		return true
	}

	nextX := x + 1
	nextY := y

	if rightmostTile {
		nextX = 0
		nextY = y + 1
	}

	if success := ta.tryVariants(nextX, nextY); success {
		return true
	}

	delete(ta.usedTileIDs, t.ID)
	// NOTE: no need to reset ta.img, as it will be overwritten later on

	return false
}

func (ta *tileAssembler) tryVariants(x, y int) bool {
	var variantsToTry []*tile.Tile
	if x > 0 {
		variantsToTry = ta.leftBorderVariants[ta.img[y][x-1].Borders.Right]
		// NOTE: will have to check top border when going through variants 1 by 1
	} else if y > 0 {
		variantsToTry = ta.topBorderVariants[ta.img[y-1][x].Borders.Bottom]
		// NOTE: x == 0, so no need to check the left border here
	}

	for _, variant := range variantsToTry {
		if ta.usedTileIDs[variant.ID] {
			continue
		}

		if x > 0 && y > 0 && !ta.img[y-1][x].MatchesBottom(variant) {
			continue
		}

		if success := ta.insertTile(variant, x, y); success {
			return true
		}
	}

	return false
}
