use std::fs::read_to_string;

use aoc_2019_08::{image::Image, part_1};

fn main() {
    let input = read_to_string("input.txt").expect("Error reading input");

    let image = Image::parse(input.trim(), 25, 6).expect("Error parsing image");

    println!("Result 1: {}", part_1(&image).expect("No layers found"));
}
