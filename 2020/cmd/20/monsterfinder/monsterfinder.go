package monsterfinder

import (
	"aoc-2020/cmd/20/positionhashmap"
	"strings"
)

func GetMonstersCount(phm positionhashmap.PositionHashMap) int {
	monsterPhm := positionhashmap.FromLines(strings.Split(`                  #
#    ##    ##    ###
 #  #  #  #  #  #   `, "\n"))

	phmVariants := phm.GetAllVariants()

	for _, phmVariant := range phmVariants {
		monstersFound := getMonstersCount(phmVariant, monsterPhm)

		if monstersFound > 0 {
			return monstersFound
		}
	}

	return 0

}

func getMonstersCount(phm positionhashmap.PositionHashMap, monsterPhm positionhashmap.PositionHashMap) int {
	monstersFound := 0

	for y := 0; y < len(phm)-len(monsterPhm); y++ {
		for x := 0; x < len(phm[0])-len(monsterPhm[0]); x++ {
			if phm.Contains(monsterPhm, x, y) {
				monstersFound++
			}
		}
	}

	return monstersFound
}
