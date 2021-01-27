use crate::{moon::Moon, position::Position, simulation::Simulation};

pub fn solve(input: &str) -> i32 {
    let moons = input
        .split("\n")
        .map(|l| l.trim())
        .filter(|l| !l.is_empty())
        .map(|l| Position::parse(l).unwrap())
        .map(Moon::new)
        .collect();
    let mut simulation = Simulation::new(moons);
    simulation.run(1000);

    simulation.total_energy()
}
