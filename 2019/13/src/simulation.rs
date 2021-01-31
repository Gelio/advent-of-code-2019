use std::{
    collections::HashMap,
    convert::{TryFrom, TryInto},
    fmt::Display,
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

impl Display for TileId {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{}",
            match self {
                TileId::Empty => ' ',
                TileId::Wall => 'W',
                TileId::Block => 'X',
                TileId::HorizontalPaddle => '-',
                TileId::Ball => 'O',
            }
        )
    }
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
    computer: Computer,
}

impl Simulation {
    pub fn new(program: Vec<isize>) -> Self {
        Self {
            computer: Computer::with_empty_input(program),
            tiles: HashMap::new(),
        }
    }

    pub fn execute(&mut self) {
        self.computer.run_till_halt();

        let tiles = &mut self.tiles;
        self.computer
            .output()
            .chunks(3)
            .for_each(|chunk| match chunk {
                [x, y, tile_id] => {
                    let tile_id: TileId = tile_id.try_into().expect("invalid tile id");
                    tiles.insert((x.clone(), y.clone()), tile_id);
                }
                _ => panic!("invalid chunk"),
            });
    }
}

impl Display for Simulation {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let min_x = self.tiles.keys().map(|(x, _)| x).min().unwrap().clone();
        let max_x = self.tiles.keys().map(|(x, _)| x).max().unwrap().clone();
        let min_y = self.tiles.keys().map(|(_, y)| y).min().unwrap().clone();
        let max_y = self.tiles.keys().map(|(_, y)| y).max().unwrap().clone();

        for y in min_y..=max_y {
            for x in min_x..=max_x {
                write!(f, "{}", self.tiles.get(&(x, y)).unwrap_or(&TileId::Empty))?;
            }
            write!(f, "\n")?;
        }

        Ok(())
    }
}
