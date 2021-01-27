use std::collections::HashSet;

use crate::{moon::parse_moons, simulation::Simulation};

pub fn solve(input: &str) -> i64 {
    let moons = parse_moons(input);
    let mut simulation = Simulation::new(moons);

    let mut encountered_situations = HashSet::new();
    encountered_situations.insert(simulation.moons.clone());

    let mut steps = 0;
    loop {
        simulation.run_single_step();
        steps += 1;

        if encountered_situations.contains(&simulation.moons) {
            break;
        }

        encountered_situations.insert(simulation.moons.clone());
    }

    steps
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
