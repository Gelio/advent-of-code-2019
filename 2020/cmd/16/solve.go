package main

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
	matchSpecFromIndex(specs, specToFieldMapping, validNearbyTickets, 0)

	return 0
}

func matchSpecFromIndex(specs []ticketFieldSpec, specToFieldMapping map[int]int, tickets [][]int, fieldIndex int) bool {
	if fieldIndex == len(specs) {
		return true
	}

	for specIndex, spec := range specs {
		if _, ok := specToFieldMapping[specIndex]; ok {
			continue
		}

		if isSpecValidForTicketsAtIndex(spec, tickets, fieldIndex) {
			// Try with this spec
			specToFieldMapping[specIndex] = fieldIndex
			if matchSpecFromIndex(specs, specToFieldMapping, tickets, fieldIndex+1) {
				return true
			}
			delete(specToFieldMapping, specIndex)
		}
	}

	return false
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
	currentTicketLoop:
		for _, fieldValue := range ticket {
			for _, spec := range specs {
				for _, fieldRange := range spec.ranges {
					if !fieldRange.Has(fieldValue) {
						break currentTicketLoop
					}
				}
			}

			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}
