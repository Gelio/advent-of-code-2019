use std::io::stdin;

use aoc_2019_07;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();

    let computer_memory: Vec<isize> =
        intcode_computer::program::parse_from_string(&buffer).unwrap();

    println!(
        "Result A: {}",
        aoc_2019_07::part_a::MaxSequenceFinder::find_max_signal(computer_memory.clone())
    );

    println!(
        "Result B: {}",
        aoc_2019_07::part_b::part_b(computer_memory.clone())
    );
}
