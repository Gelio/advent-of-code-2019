package main

import "fmt"

var targetBag = "shiny gold"

func SolveA(lines []string) (visitedBagsCount int, err error) {
	// Edge A -> B in the graph means that bag A can held inside bag B
	bagGraph := make(map[string][]string)

	for _, line := range lines {
		var rule rule
		rule, err = parseLine(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, nestedBag := range rule.contents {
			bagGraph[nestedBag.color] = append(bagGraph[nestedBag.color], rule.bagColor)
		}
	}

	openBags := make([]string, 1, len(bagGraph))
	openBags[0] = targetBag
	closedBags := make(map[string]bool)

	for len(openBags) > 0 {
		bagToVisit := openBags[0]
		openBags = openBags[1:]

		wrappingBags, ok := bagGraph[bagToVisit]
		if !ok {
			continue
		}

		for _, wrappingBag := range wrappingBags {
			if bagAlreadyVisited := closedBags[wrappingBag]; bagAlreadyVisited {
				continue
			}

			openBags = append(openBags, wrappingBag)
			closedBags[wrappingBag] = true
			visitedBagsCount++
		}
	}

	return
}
