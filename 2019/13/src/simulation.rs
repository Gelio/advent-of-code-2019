use std::{
    collections::HashMap,
    convert::{TryFrom, TryInto},
};

use intcode_computer::Computer;

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum TileId {
    Empty,
    Wall,
    Block,
    HorizontalPaddle,
    Ball,
}

impl TryFrom<&isize> for TileId {
    type Error = String;

    fn try_from(value: &isize) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Self::Empty),
            1 => Ok(Self::Wall),
            2 => Ok(Self::Block),
            3 => Ok(Self::HorizontalPaddle),
            4 => Ok(Self::Ball),
            _ => Err(format!("invalid tile id {}", value)),
        }
    }
}

pub struct Simulation {
    pub tiles: HashMap<(isize, isize), TileId>,
}

impl Simulation {
    pub fn execute(program: Vec<isize>) -> Self {
        let mut computer = Computer::with_empty_input(program);
        computer.run_till_halt();

        let mut tiles = HashMap::new();

        let output = computer.output();
        output.chunks(3).for_each(|chunk| match chunk {
            [x, y, tile_id] => {
                let tile_id: TileId = tile_id.try_into().expect("invalid tile id");
                tiles.insert((x.clone(), y.clone()), tile_id);
            }
            _ => panic!("invalid chunk"),
        });

        Self { tiles }
    }
}
