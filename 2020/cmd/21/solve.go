package main

import (
	"sort"
	"strings"
)

// ingredient contains 0 or 1 allergens
// allergen is in 1 ingredient

func solveA(foods []food) int {
	as := newAssigner()
	as.FindAssignment(foods)

	ingredientCounts := getIngredientsCount(foods)

	ingredientsAssigned := make(map[ingredientName]bool)
	for _, ingredient := range as.allergenAssignments {
		ingredientsAssigned[ingredient] = true
	}

	sum := 0
	for ingredient, count := range ingredientCounts {
		if !ingredientsAssigned[ingredient] {
			sum += count
		}
	}

	return sum
}

func solveB(foods []food) string {
	as := newAssigner()
	as.FindAssignment(foods)

	var allergens []string
	for allergen := range as.allergenAssignments {
		allergens = append(allergens, string(allergen))
	}
	sort.Strings(allergens)

	var ingredientsWithAllergens []string
	for _, allergen := range allergens {
		ingredient := as.allergenAssignments[allergenName(allergen)]
		ingredientsWithAllergens = append(ingredientsWithAllergens, string(ingredient))
	}

	return strings.Join(ingredientsWithAllergens, ",")
}

type assigner struct {
	possibleIngredients map[allergenName]map[ingredientName]bool
	allergenAssignments map[allergenName]ingredientName
}

func newAssigner() assigner {
	var as assigner
	as.possibleIngredients = make(map[allergenName]map[ingredientName]bool)
	as.allergenAssignments = make(map[allergenName]ingredientName)

	return as
}

func (as *assigner) FindAssignment(foods []food) {
	for _, f := range foods {
		as.addFood(f)
	}

	as.assignSinglePossibilities()

	// NOTE: yay, there is only 1 possible assignment If there were many, one would have to do
	// backtracking on the possibleIngredients for each allergen
}

func (as *assigner) addFood(f food) {
	for _, allergen := range f.allegrens {
		ingredientsInCurrentFood := make(map[ingredientName]bool)
		for _, ingredient := range f.ingredients {
			ingredientsInCurrentFood[ingredient] = true
		}

		if possibleIngredients, ok := as.possibleIngredients[allergen]; ok {
			for ingredient := range possibleIngredients {
				if !ingredientsInCurrentFood[ingredient] {
					delete(possibleIngredients, ingredient)
				}
			}
		} else {
			as.possibleIngredients[allergen] = ingredientsInCurrentFood
		}
	}
}

// Assign those allergens that only have 1 possible ingredient
func (as *assigner) assignSinglePossibilities() {
	assignedSome := true
	for assignedSome {
		assignedSome = false

		for allergen, ingredients := range as.possibleIngredients {
			if len(ingredients) != 1 {
				continue
			}

			var onlyIngredient ingredientName
			for ingredient := range ingredients {
				onlyIngredient = ingredient
			}

			assignedSome = true
			as.allergenAssignments[allergen] = onlyIngredient
			delete(as.possibleIngredients, allergen)

			for _, otherIngredients := range as.possibleIngredients {
				delete(otherIngredients, onlyIngredient)
			}
		}
	}
}

func getIngredientsCount(foods []food) map[ingredientName]int {
	counts := make(map[ingredientName]int)

	for _, f := range foods {
		for _, ingredient := range f.ingredients {
			counts[ingredient]++
		}
	}

	return counts
}
