package main

import (
	"strconv"
	"strings"
)

func solveA(specs []ticketFieldSpec, nearbyTickets [][]int) int {
	scanningErrorRate := 0

	for _, ticket := range nearbyTickets {
		for _, fieldValue := range ticket {
			valid := false
		currentValueLoop:
			for _, spec := range specs {
				for _, fieldRange := range spec.ranges {
					if fieldRange.Has(fieldValue) {
						valid = true
						break currentValueLoop
					}
				}
			}

			if !valid {
				scanningErrorRate += fieldValue
			}
		}
	}

	return scanningErrorRate
}

func parseSpecs(lines []string) ([]ticketFieldSpec, error) {
	var specs []ticketFieldSpec

	for _, line := range lines {
		spec, err := NewTicketFieldSpec(line)
		if err != nil {
			return nil, err
		}

		specs = append(specs, spec)
	}

	return specs, nil
}

func parseTickets(lines []string) ([][]int, error) {
	var tickets [][]int

	for _, line := range lines {
		var ticketNums []int
		for _, num := range strings.Split(line, ",") {
			parsedNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			ticketNums = append(ticketNums, parsedNum)
		}

		tickets = append(tickets, ticketNums)
	}

	return tickets, nil
}
