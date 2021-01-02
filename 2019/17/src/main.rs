use aoc_2019_17::part_1;
use intcode_computer::program;
use std::fs::read_to_string;

fn main() {
    let program =
        program::parse_from_string(&read_to_string("input.txt").expect("Cannot read input file"))
            .unwrap();

    println!("Result A: {}", part_1(program));
}
