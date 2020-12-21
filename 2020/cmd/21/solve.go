package main

// ingredient contains 0 or 1 allergens
// allergen is in 1 ingredient

func solveA(foods []food) (int, error) {
	as := newAssigner(foods)

	as.assignAllergens(0, 0)

	ingredientCounts := getIngredientsCount(foods)

	sum := 0
	for ingredient, count := range ingredientCounts {
		if !as.ingredientEverAssigned[ingredient] {
			sum += count
		}
	}

	return sum, nil
}

type assigner struct {
	allergensAssignment    map[allergenName]ingredientName
	ingredientsAssigned    map[ingredientName]bool
	ingredientEverAssigned map[ingredientName]bool
	foods                  []food
}

func newAssigner(foods []food) assigner {
	var as assigner
	as.allergensAssignment = make(map[allergenName]ingredientName)
	as.ingredientsAssigned = make(map[ingredientName]bool)
	as.ingredientEverAssigned = make(map[ingredientName]bool)
	as.foods = foods

	return as
}

func (as *assigner) assignAllergens(foodIndex, allergenIndex int) bool {
	if foodIndex == len(as.foods) {
		for ingredient := range as.ingredientsAssigned {
			as.ingredientEverAssigned[ingredient] = true
		}

		return true
	}

	if allergenIndex == len(as.foods[foodIndex].allegrens) {
		return as.assignAllergens(foodIndex+1, 0)
	}

	allergen := as.foods[foodIndex].allegrens[allergenIndex]
	f := as.foods[foodIndex]

	if assignedIngredient, allergenAssigned := as.allergensAssignment[allergen]; allergenAssigned {
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

		return as.assignAllergens(foodIndex, allergenIndex+1)
	}

	// unassignedAllergen := true

	// Match unassigned allergen to an unassigned ingredient
	for _, ingredient := range f.ingredients {
		if as.ingredientsAssigned[ingredient] {
			continue
		}

		as.allergensAssignment[allergen] = ingredient
		as.ingredientsAssigned[ingredient] = true

		if success := as.assignAllergens(foodIndex, allergenIndex+1); success {
			return true
			// unassignedAllergen = false
		}

		delete(as.allergensAssignment, allergen)
		delete(as.ingredientsAssigned, ingredient)
	}

	return false
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
