use nom::combinator::all_consuming;
use simulation::{parser::map_with_moves, Simulation};

mod map;
mod map_tiles;
mod moves;
mod player;
mod simulation;

fn main() {
    let input = include_str!("../input.txt").trim_end();
    let (_, (map, moves)) = all_consuming(map_with_moves)(input).expect("Parsing error");
    let mut simulation = Simulation::new(map);
    for m in moves {
        simulation.simulate_move(m)
    }

    println!("Part 1: {0}", simulation.player().password());
}
