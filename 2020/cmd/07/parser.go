package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// QUESTION: would you rename anything here? e.g. shorten some variable names?

func parseLine(line string) (r rule, err error) {
	bagRegexpExpression := "(\\w+) (\\w+) bags?"

	{
		result := regexp.MustCompile(fmt.Sprintf("^%s", bagRegexpExpression)).FindStringSubmatch(line)
		if len(result) == 0 {
			err = fmt.Errorf("Cannot find main bag in %#v", line)
			return
		}

		r.bagColor = fmt.Sprintf("%s %s", result[1], result[2])
	}

	{
		var matched bool
		matched, err = regexp.MatchString("no other bags.", line)
		if err != nil {
			return
		}

		if matched {
			return
		}
	}

	bagWithQuantityRegexp := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, bagRegexpExpression))

	contentsMatches := bagWithQuantityRegexp.FindAllStringSubmatch(line, -1)

	if len(contentsMatches) == 0 {
		err = fmt.Errorf("Cannot match bag contents: %s", line)
		return
	}

	for i, contentsBag := range contentsMatches {
		quantityStr := contentsBag[1]

		if quantityStr == "" {
			break
		}

		var quantity int
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			err = fmt.Errorf("Cannot parse quantity %#v at index %d: %v", quantityStr, i, err)
			return
		}

		bag := &bagWithQuantity{
			quantity: quantity,
			color:    fmt.Sprintf("%s %s", contentsBag[2], contentsBag[3]),
		}
		r.contents = append(r.contents, bag)
	}

	return
}
