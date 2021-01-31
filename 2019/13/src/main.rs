use std::fs::read_to_string;

use aoc_2019_13::part_a;
use intcode_computer::program::parse_from_string;

fn main() {
    let input = read_to_string("input.txt").expect("cannot read input");
    let program = parse_from_string(&input).expect("cannot parse program");

    println!("Result A: {}", part_a::solve(program));
}
