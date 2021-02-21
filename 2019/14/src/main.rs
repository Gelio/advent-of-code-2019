use std::fs::read_to_string;

use aoc_2019_14::{production_graph::ProductionGraph, rules::parse_production_rules};

fn main() {
    let input = read_to_string("input.txt").expect("cannot read input file");
    let parsing_result = parse_production_rules(
        &input
            .split("\n")
            .map(str::trim)
            .filter(|s| !s.is_empty())
            .collect::<Vec<_>>(),
    )
    .expect("cannot parse rules");
    let graph = ProductionGraph::new(parsing_result.rules);

    println!(
        "Result A: {}",
        graph.find_path(parsing_result.fuel_index, parsing_result.ore_index)
    );
}
