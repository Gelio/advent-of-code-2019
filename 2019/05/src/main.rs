use std::fs::read_to_string;

use aoc_2019_05::Computer;

fn main() {
    let input = read_to_string("input.txt").unwrap();

    let memory: Vec<isize> = input
        .trim()
        .split(",")
        .map(|x| x.parse().expect(&format!("Cannot parse number {}", x)))
        .collect();

    let mut computer = Computer::new(memory.clone(), vec![1]);
    computer.run();

    println!("Result A: {:?}", computer.output);

    let mut computer = Computer::new(memory, vec![5]);
    computer.run();

    println!("Result B: {:?}", computer.output);
}
