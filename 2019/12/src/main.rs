use std::fs::read_to_string;

use aoc_2019_12::{part_a, part_b};

fn main() {
    let input = read_to_string("input.txt").expect("cannot read input");

    println!("Result A: {}", part_a::solve(&input));
    println!("Result B: {}", part_b::solve(&input));
}
