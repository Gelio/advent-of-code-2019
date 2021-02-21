use std::collections::HashMap;

use crate::rules::ProductionRule;

pub struct ProductionGraph {
    /// Contains rules that produce index
    edges: HashMap<usize, Vec<ProductionRule>>,
}

impl ProductionGraph {
    pub fn new(rules: Vec<ProductionRule>) -> Self {
        let mut edges: HashMap<usize, Vec<ProductionRule>> = HashMap::new();

        for rule in rules {
            edges.entry(rule.output.chemical).or_default().push(rule)
        }

        Self { edges }
    }

    pub fn find_path(&self, fuel: usize, ore: usize) -> usize {
        let mut needed_ore = 0;
        let mut pending_ingredients: HashMap<usize, usize> = HashMap::new();
        let mut leftover_ingredients: HashMap<usize, usize> = HashMap::new();

        pending_ingredients.insert(fuel, 1);

        while !pending_ingredients.is_empty() {
            let ingredient = *pending_ingredients.keys().next().unwrap();
            let mut quantity = pending_ingredients.remove(&ingredient).unwrap();

            leftover_ingredients
                .get_mut(&ingredient)
                .map(|leftover_quantity| {
                    let from_leftovers = quantity.min(*leftover_quantity);
                    quantity -= from_leftovers;
                    *leftover_quantity -= from_leftovers;
                });

            if quantity == 0 {
                continue;
            }

            let rule = self.edges.get(&ingredient).unwrap().first().expect(
                "assumption that there is only 1 production rule with that output is invalid",
            );
            let mut times_to_apply_rule = quantity / rule.output.quantity;
            if times_to_apply_rule * rule.output.quantity != quantity {
                times_to_apply_rule += 1;
            }
            *leftover_ingredients.entry(ingredient).or_default() +=
                times_to_apply_rule * rule.output.quantity - quantity;

            (&rule.ingredients).into_iter().for_each(|item| {
                let quantity_to_add = item.quantity * times_to_apply_rule;
                if item.chemical == ore {
                    needed_ore += quantity_to_add;
                } else {
                    *pending_ingredients.entry(item.chemical).or_default() += quantity_to_add;
                }
            })
        }

        needed_ore
    }
}
