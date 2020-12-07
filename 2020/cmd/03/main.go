package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type slope struct {
	right, down int
}

const treeCharacter = '#'

func main() {
	reader := bufio.NewReader(os.Stdin)

	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}
	var problemMap []string
	var lineWrappingLength int

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("Error while reading input", err)
			return
		}

		line = strings.TrimSpace(line)
		lineLength := len(line)
		if lineLength == 0 {
			break
		}
		problemMap = append(problemMap, line)

		if lineWrappingLength == 0 {
			lineWrappingLength = lineLength
		}
	}

	result := 1

	for _, slope := range slopes {
		result *= countTreesEncountered(&slope, &problemMap)
	}

	fmt.Println("Result: ", result)
}

func countTreesEncountered(slope *slope, problemMap *[]string) int {
	treesEncountered := 0
	positionX := 0

	for y := 0; y < len(*problemMap); y += (*slope).down {
		line := (*problemMap)[y]
		characterEncountered := line[positionX]

		if characterEncountered == treeCharacter {
			treesEncountered++
		}

		positionX = (positionX + (*slope).right) % len(line)
	}

	return treesEncountered

}
