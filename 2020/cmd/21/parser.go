package main

import (
	"fmt"
	"regexp"
	"strings"
)

type ingredientName string
type allergenName string

type food struct {
	ingredients []ingredientName
	allegrens   []allergenName
}

func parseFood(line string) (food, error) {
	lineRegexp, err := regexp.Compile(`^(.*) \(contains (.*)\)$`)
	if err != nil {
		return food{}, fmt.Errorf("cannot compile line regexp: %w", err)
	}

	matches := lineRegexp.FindStringSubmatch(line)
	if len(matches) != 3 {
		return food{}, fmt.Errorf("cannot match line %q", line)
	}

	var f food
	for _, ingredient := range strings.Split(matches[1], " ") {
		f.ingredients = append(f.ingredients, ingredientName(ingredient))
	}
	for _, allergen := range strings.Split(matches[2], ", ") {
		f.allegrens = append(f.allegrens, allergenName(allergen))
	}

	return f, nil
}
