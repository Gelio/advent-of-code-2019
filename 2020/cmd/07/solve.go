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

type bagGraph map[string][](*bagWithQuantity)

func SolveB(lines []string) (bagsRequired int, err error) {
	g := make(bagGraph)

	for _, line := range lines {
		var rule rule
		rule, err = parseLine(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		g[rule.bagColor] = make([](*bagWithQuantity), 0, len(rule.contents))

		for _, nestedBag := range rule.contents {
			g[rule.bagColor] = append(g[rule.bagColor], nestedBag)
		}
	}

	allBags, err := g.countNestedBagWithMainBag(targetBag)
	bagsRequired = allBags - 1

	return
}

func (g bagGraph) countNestedBagWithMainBag(bagColor string) (result int, err error) {
	nestedBags, ok := g[bagColor]

	if !ok {
		err = fmt.Errorf("Unknown bag color encountered: %s", bagColor)
		return
	}

	result = 1

	for _, bag := range nestedBags {
		var r int
		r, err = g.countNestedBagWithMainBag(bag.color)

		if err != nil {
			return
		}

		result += r * bag.quantity
	}

	return
}
