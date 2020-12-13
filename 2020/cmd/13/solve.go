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

func solveB(startTimestamp int, busIDs []int) int {
	firstBusDepartureTimestamp := (startTimestamp/busIDs[0] + 1) * busIDs[0]

	var actualBusIndexes []int

	for i, busID := range busIDs {
		if busID != cross {
			actualBusIndexes = append(actualBusIndexes, i)
		}
	}

	for t := firstBusDepartureTimestamp; true; t += busIDs[0] {
		allMatch := true

		for _, busIndex := range actualBusIndexes {
			if busDeparts := (t+busIndex)%busIDs[busIndex] == 0; !busDeparts {
				allMatch = false
				break
			}
		}

		if allMatch {
			return t
		}
	}

	return -1
}
