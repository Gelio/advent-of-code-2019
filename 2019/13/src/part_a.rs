use crate::simulation::{Simulation, TileId};

pub fn solve(program: Vec<isize>) -> usize {
    let mut simulation = Simulation::new(program);
    simulation.execute();

    simulation
        .tiles
        .into_iter()
        .filter(|(_, tile_id)| tile_id == &TileId::Block)
        .count()
}
