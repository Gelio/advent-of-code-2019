use std::fs::read_to_string;

use aoc_2019_08::{image::Image, part_1, part_2};

fn main() {
    let input = read_to_string("input.txt").expect("Error reading input");

    let image = Image::parse(input.trim(), 25, 6).expect("Error parsing image");

    println!(
        "Result 1: {}",
        part_1(&image).expect("Error in part 1: no layers found")
    );
    println!(
        "Result 2:\n{}",
        part_2(&image).expect("Error in part 2: could not compose layers")
    )
}
