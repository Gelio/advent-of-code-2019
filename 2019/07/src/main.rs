use std::io::stdin;

use aoc_2019_07;

fn main() {
    let mut buffer = String::new();
    stdin().read_line(&mut buffer).unwrap();

    let computer_memory: Vec<isize> = buffer
        .trim()
        .split(",")
        .map(|c| c.parse().expect(&format!("Cannot parse {}", c).to_owned()))
        .collect();

    println!(
        "Result A: {}",
        aoc_2019_07::part_a::MaxSequenceFinder::find_max_signal(computer_memory)
    );
}
