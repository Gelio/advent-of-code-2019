use crate::mixer::{part_1, part_2};

mod linked_list;
mod mixer;

fn main() {
    let numbers: Vec<_> = include_str!("./input.txt")
        .split_whitespace()
        .map(str::parse)
        .collect::<Result<_, _>>()
        .expect("Could not parse some number");

    println!("Part 1: {0}", part_1(numbers.clone()));
    println!("Part 2: {0}", part_2(numbers));
}
