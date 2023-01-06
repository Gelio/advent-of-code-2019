use crate::mixer::part_1;

mod linked_list;
mod mixer;

fn main() {
    let numbers: Vec<i32> = include_str!("./input.txt")
        .split_whitespace()
        .map(str::parse)
        .collect::<Result<_, _>>()
        .expect("Could not parse some number");

    println!("Part A: {0}", part_1(numbers));
}
