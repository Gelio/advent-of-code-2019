use std::fs::read_to_string;

use aoc_2019_06::OrbitsGraph;

fn main() {
    let input = read_to_string("input.txt").expect("Cannot read input file");

    let orbits_graph = OrbitsGraph::parse(input.trim().lines()).expect("Cannot parse orbits");

    println!("Result A: {}", orbits_graph.get_total_orbits("COM", 0));
}
