use std::fs::read_to_string;

use aoc_2019_11::{
    map_printer,
    move_emulator::{Color, MoveEmulator},
    point::Point,
};
use intcode_computer::program;

fn main() {
    let input = read_to_string("input.txt").expect("cannot read input");
    let program = program::parse_from_string(&input).expect("cannot parse program");
    let mut emulator = MoveEmulator::new(program.clone());

    emulator.run_till_halt();

    println!("Result A: {}", emulator.map.len());

    let mut emulator = MoveEmulator::new(program);
    emulator.map.insert(Point::default(), Color::White);
    emulator.run_till_halt();

    println!("Result B:");
    map_printer::print_map(&emulator.map);
}
