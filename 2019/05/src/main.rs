use std::fs::read_to_string;

use aoc_2019_05::Computer;

fn main() {
    let input = read_to_string("input.txt").unwrap();

    let memory: Vec<isize> = input
        .trim()
        .split(",")
        .map(|x| x.parse().expect(&format!("Cannot parse number {}", x)))
        .collect();

    let input = vec![1];

    let mut computer = Computer::new(memory, input);
    computer.run();

    println!("Result: {:?}", computer.output);
}
