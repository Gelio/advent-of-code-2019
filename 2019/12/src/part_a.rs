use crate::{moon::parse_moons, simulation::Simulation};

pub fn solve(input: &str) -> i32 {
    let moons = parse_moons(input);
    let mut simulation = Simulation::new(moons);
    simulation.run(1000);

    simulation.total_energy()
}
