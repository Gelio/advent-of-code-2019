package main

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
