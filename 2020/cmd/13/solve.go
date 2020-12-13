package main

import (
	"strconv"
	"strings"
)

func parseBusIDs(line string) ([]int, error) {
	stringifiedIDs := strings.Split(line, ",")

	var busIDs []int

	for _, stringifiedID := range stringifiedIDs {
		if stringifiedID == "x" {
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
		busTrips := timestamp/busID + 1
		timeToWait := busTrips*busID - timestamp
		if solution == nil || solution.timeToWait > timeToWait {
			solution = &solutionA{timeToWait, busID}
		}
	}

	return solution.busID * solution.timeToWait
}
