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
