use std::fs::read_to_string;

use aoc_2019_06::{get_total_orbits, parse_orbits};

fn main() {
    let input = read_to_string("input.txt").expect("Cannot read input file");

    let orbits_graph = parse_orbits(input.trim().lines()).expect("Cannot parse orbits");

    println!("Result A: {}", get_total_orbits(&orbits_graph, "COM", 0));
}
