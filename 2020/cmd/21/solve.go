package main

// ingredient contains 0 or 1 allergens
// allergen is in 1 ingredient

// TODO: add a struct instead of having separate methods

func solveA(foods []food) (int, error) {
	allergensAssignment := make(map[allergenName]ingredientName)
	ingredientsAssigned := make(map[ingredientName]bool)
	ingredientEverAssigned := make(map[ingredientName]bool)

	assignAllergenRecursive(allergensAssignment, ingredientsAssigned, ingredientEverAssigned, foods)

	ingredientCounts := getIngredientsCount(foods)

	sum := 0
	for ingredient, count := range ingredientCounts {
		if !ingredientEverAssigned[ingredient] {
			sum += count
		}
	}

	return sum, nil
}

func assignAllergenRecursive(allergensAssignment map[allergenName]ingredientName, ingredientsAssigned map[ingredientName]bool, ingredientEverAssigned map[ingredientName]bool, foods []food) bool {
	if len(foods) == 0 {
		for ingredient := range ingredientsAssigned {
			ingredientEverAssigned[ingredient] = true
		}

		return true
	}

	f := foods[0]
	unassignedAllergen := false

	// TODO: try to assign a single allergen, and then assign other allergens recursively.
	// After assigning all alergens, move to the next rule
	for _, allergen := range f.allegrens {
		if assignedIngredient, allergenAssigned := allergensAssignment[allergen]; allergenAssigned {
			// Check if the assigned ingredient appears in the ingredients list
			hasIngredient := false
			for _, ingredient := range f.ingredients {
				if ingredient == assignedIngredient {
					hasIngredient = true
					break
				}
			}
			if !hasIngredient {
				return false
			}
			continue
		}

		unassignedAllergen = true

		// Match unassigned allergen to an unassigned ingredient
		for _, ingredient := range f.ingredients {
			if ingredientsAssigned[ingredient] {
				continue
			}

			allergensAssignment[allergen] = ingredient
			ingredientsAssigned[ingredient] = true

			if success := assignAllergenRecursive(allergensAssignment, ingredientsAssigned, ingredientEverAssigned, foods[1:]); success {
				unassignedAllergen = false
			}

			delete(allergensAssignment, allergen)
			delete(ingredientsAssigned, ingredient)
		}

		if unassignedAllergen {
			return false
		}
	}

	return assignAllergenRecursive(allergensAssignment, ingredientsAssigned, ingredientEverAssigned, foods[1:])
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
