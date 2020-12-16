package main

import (
	"fmt"
	"strings"
)

func solveA(specs []ticketFieldSpec, nearbyTickets [][]int) int {
	scanningErrorRate := 0

	for _, ticket := range nearbyTickets {
		for _, fieldValue := range ticket {
			valid := false
			for _, spec := range specs {
				if spec.Matches(fieldValue) {
					valid = true
					break
				}
			}

			if !valid {
				scanningErrorRate += fieldValue
			}
		}
	}

	return scanningErrorRate
}

func solveB(specs []ticketFieldSpec, myTicket []int, nearbyTickets [][]int) int {
	validNearbyTickets := getValidTickets(specs, nearbyTickets)

	// specToFieldMapping[specIndex] = fieldIndex
	specToFieldMapping := make(map[int]int)
	matchSpecFromIndex(specs, specToFieldMapping, validNearbyTickets)

	result := 1
	for i, spec := range specs {
		if strings.HasPrefix(spec.name, "departure") {
			result *= myTicket[specToFieldMapping[i]]
		}
	}

	return result
}

func matchSpecFromIndex(specs []ticketFieldSpec, specToFieldMapping map[int]int, tickets [][]int) {
	var fieldMatched []bool
	for range specs {
		fieldMatched = append(fieldMatched, false)
	}

	for fieldIndex, specFieldsMatched := 0, 0; specFieldsMatched < len(fieldMatched); fieldIndex = (fieldIndex + 1) % len(fieldMatched) {
		if fieldMatched[fieldIndex] {
			continue
		}

		specIndex, err := getOnlyMatchingSpecIndexForFieldIndex(specs, specToFieldMapping, tickets, fieldIndex)
		if err != nil {
			continue
		}

		specFieldsMatched++
		specToFieldMapping[specIndex] = fieldIndex
		fieldMatched[fieldIndex] = true
	}
}

func getOnlyMatchingSpecIndexForFieldIndex(specs []ticketFieldSpec, specToFieldMapping map[int]int, tickets [][]int, fieldIndex int) (int, error) {
	var foundSpecIndex int
	specFound := false

	for specIndex, spec := range specs {
		if _, ok := specToFieldMapping[specIndex]; ok {
			continue
		}

		if isSpecValidForTicketsAtIndex(spec, tickets, fieldIndex) {
			if specFound {
				return 0, fmt.Errorf("found 2 or more matching specs for field at index %d", fieldIndex)
			}

			foundSpecIndex = specIndex
			specFound = true
		}
	}

	return foundSpecIndex, nil
}

func isSpecValidForTicketsAtIndex(spec ticketFieldSpec, tickets [][]int, index int) bool {
	for _, ticket := range tickets {
		if !spec.Matches(ticket[index]) {
			return false
		}
	}

	return true
}

func getValidTickets(specs []ticketFieldSpec, nearbyTickets [][]int) [][]int {
	var validTickets [][]int

	for _, ticket := range nearbyTickets {
		ticketValid := true

		for _, fieldValue := range ticket {
			someSpecMatches := false

			for _, spec := range specs {
				if spec.Matches(fieldValue) {
					someSpecMatches = true
					break
				}
			}

			if !someSpecMatches {
				ticketValid = false
				break
			}
		}

		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}
