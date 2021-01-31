use crate::simulation::{Simulation, TileId};

pub fn solve(program: Vec<isize>) -> usize {
    let simulation = Simulation::execute(program);

    simulation
        .tiles
        .into_iter()
        .filter(|(_, tile_id)| tile_id == &TileId::Block)
        .count()
}
