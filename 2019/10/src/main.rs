use std::fs::read_to_string;

use aoc_2019_10::{map::AsteroidMap, part_a, part_b::find_200th_removed_asteroid};

fn main() {
    let input = read_to_string("input.txt").expect("Cannot read input");
    let mut map =
        AsteroidMap::parse(&input.split_whitespace().collect()).expect("Cannot parse map");

    let best_asteroid = part_a::get_best_asteroid(&map).expect("cannot find best asteroid");
    println!("Result A: {}", best_asteroid.visible_asteroids.len());

    let target_asteroid = find_200th_removed_asteroid(&mut map, &best_asteroid.pos);
    println!("Result B: {}", target_asteroid.x * 100 + target_asteroid.y);
}
