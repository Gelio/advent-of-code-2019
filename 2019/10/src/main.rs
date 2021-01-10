use std::fs::read_to_string;

use aoc_2019_10::{map::AsteroidMap, part_a};

fn main() {
    let input = read_to_string("input.txt").expect("Cannot read input");
    let map = AsteroidMap::parse(&input.split_whitespace().collect()).expect("Cannot parse map");

    println!(
        "Result A: {}",
        part_a::get_best_asteroid_neighbors_count(&map)
    );
}
