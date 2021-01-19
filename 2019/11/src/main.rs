use std::fs::read_to_string;

use aoc_2019_11::move_emulator::MoveEmulator;
use intcode_computer::program;

fn main() {
    let input = read_to_string("input.txt").expect("cannot read input");
    let program = program::parse_from_string(&input).expect("cannot parse program");
    let mut emulator = MoveEmulator::new(program);

    emulator.run_till_halt();

    println!("Result A: {}", emulator.map.len());
}
