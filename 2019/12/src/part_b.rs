use std::{
    collections::{hash_map::DefaultHasher, HashSet},
    hash::{Hash, Hasher},
};

use crate::{
    moon::{parse_moons, Moon},
    position::Position,
    simulation::Simulation,
};

pub fn solve(input: &str) -> u64 {
    let moons = parse_moons(input);
    let mut simulation = Simulation::new(moons);
    let mut x_cycle_finder = CycleLengthFinder::new(get_x_dimension);
    let mut y_cycle_finder = CycleLengthFinder::new(get_y_dimension);
    let mut z_cycle_finder = CycleLengthFinder::new(get_z_dimension);

    loop {
        match (
            x_cycle_finder.cycle_length(),
            y_cycle_finder.cycle_length(),
            z_cycle_finder.cycle_length(),
        ) {
            (Some(x), Some(y), Some(z)) => {
                // NOTE: subtracting 1 from each cycle length, as the initial state registration
                // increments the counter by 1
                return lcm(x - 1, lcm(y - 1, z - 1));
            }
            _ => {
                let moons = &simulation.moons;
                x_cycle_finder.register_new_state(moons);
                y_cycle_finder.register_new_state(moons);
                z_cycle_finder.register_new_state(moons);

                simulation.run_single_step();
            }
        }
    }
}

// Lowest common multiple
// See http://www.programming-algorithms.net/article/42865/Least-common-multiple
fn lcm(a: u64, b: u64) -> u64 {
    a * b / gcd(a, b)
}

// Greatest common denominator
fn gcd(mut a: u64, mut b: u64) -> u64 {
    let mut remainder: u64;

    while b != 0 {
        remainder = a % b;
        a = b;
        b = remainder;
    }

    a
}

fn get_x_dimension(pos: &Position) -> i32 {
    pos.x
}
fn get_y_dimension(pos: &Position) -> i32 {
    pos.y
}
fn get_z_dimension(pos: &Position) -> i32 {
    pos.z
}

#[derive(Debug)]
struct CycleLengthFinder<G> {
    encountered: HashSet<u64>,
    found: bool,
    cycle_length: u64,
    dimention_getter: G,
}

impl<G> CycleLengthFinder<G>
where
    G: Fn(&Position) -> i32,
{
    pub fn new(dimention_getter: G) -> Self {
        Self {
            dimention_getter,
            encountered: HashSet::new(),
            found: false,
            cycle_length: 0,
        }
    }

    pub fn register_new_state(&mut self, state: &Vec<Moon>) {
        if self.found {
            return;
        }

        self.cycle_length += 1;

        let mut hasher = DefaultHasher::new();
        let get_dimension = &self.dimention_getter;

        for moon in state.iter() {
            let position = get_dimension(moon.position());
            let velocity = get_dimension(moon.velocity());
            position.hash(&mut hasher);
            velocity.hash(&mut hasher);
        }

        let state_hash = hasher.finish();
        if self.encountered.contains(&state_hash) {
            self.found = true;
            return;
        }

        self.encountered.insert(state_hash);
    }

    pub fn cycle_length(&self) -> Option<u64> {
        if self.found {
            Some(self.cycle_length)
        } else {
            None
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn works() {
        let input = "<x=-1, y=0, z=2>
            <x=2, y=-10, z=-7>
            <x=4, y=-8, z=8>
            <x=3, y=5, z=-1>";
        let result = solve(input);

        assert_eq!(result, 2772);
    }

    #[test]
    fn works_reasonably_fast() {
        let input = "<x=-8, y=-10, z=0>
            <x=5, y=5, z=10>
            <x=2, y=-7, z=3>
            <x=9, y=-8, z=-3>";
        let result = solve(input);

        assert_eq!(result, 4686774924);
    }
}
