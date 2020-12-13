package main

import (
	"strconv"
	"strings"
)

// Integer value for "x" in the input
const cross = -1

func parseBusIDs(line string) ([]int, error) {
	stringifiedIDs := strings.Split(line, ",")

	var busIDs []int

	for _, stringifiedID := range stringifiedIDs {
		if stringifiedID == "x" {
			busIDs = append(busIDs, cross)
			continue
		}

		id, err := strconv.Atoi(stringifiedID)
		if err != nil {
			return nil, err
		}

		busIDs = append(busIDs, id)
	}

	return busIDs, nil
}

type solutionA struct {
	timeToWait, busID int
}

func solveA(timestamp int, busIDs []int) int {
	var solution *solutionA

	for _, busID := range busIDs {
		if busID == cross {
			continue
		}

		busTrips := timestamp/busID + 1
		timeToWait := busTrips*busID - timestamp
		if solution == nil || solution.timeToWait > timeToWait {
			solution = &solutionA{timeToWait, busID}
		}
	}

	return solution.busID * solution.timeToWait
}

func solveB(busIDs []int) int {
	// Solved based on the example in Polish Chinese reminder theorem Wikipedia article
	// https://pl.wikipedia.org/wiki/Chi%C5%84skie_twierdzenie_o_resztach#Przyk%C5%82ad
	var a, b int

	for i, busID := range busIDs {
		if busID == cross {
			continue
		}

		if i == 0 {
			b = busID
			continue
		}

		for x := a; true; x += b {
			if (x+i)%busID == 0 {
				a = x
				b *= busID
				break
			}
		}
	}

	return a
}
