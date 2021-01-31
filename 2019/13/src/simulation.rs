use std::{
    cell::RefCell,
    collections::HashMap,
    convert::{TryFrom, TryInto},
    fmt::Display,
    rc::Rc,
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
                TileId::HorizontalPaddle => '#',
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
    input: Rc<RefCell<Vec<isize>>>,
    pub score: isize,
    pub computer: Computer,
}

impl Simulation {
    pub fn new(program: Vec<isize>) -> Self {
        let input = Rc::new(RefCell::new(Vec::new()));
        Self {
            input: Rc::clone(&input),
            computer: Computer::new(program, input),
            tiles: HashMap::new(),
            score: 0,
        }
    }

    pub fn execute(&mut self) {
        let mut output_instruction_id = 0;
        loop {
            match self.computer.parse_and_exec_once() {
                intcode_computer::Instruction::WriteOutput { val: -1 } if self.tiles.is_empty() => {
                    // Write 2 more pieces of the "score" update
                    self.computer.parse_and_exec_once();
                    self.computer.parse_and_exec_once();
                    break;
                }
                intcode_computer::Instruction::WriteOutput { val: 4 }
                    if output_instruction_id % 3 == 2 && !self.tiles.is_empty() =>
                {
                    // Ball printed
                    break;
                }
                intcode_computer::Instruction::Halt => {
                    break;
                }
                intcode_computer::Instruction::WriteOutput { .. } => {
                    output_instruction_id += 1;
                }
                _ => {}
            }
        }

        let tiles = &mut self.tiles;
        let score = &mut self.score;
        self.computer
            .output()
            .chunks(3)
            .for_each(|chunk| match chunk {
                [-1, 0, s] => {
                    *score = s.clone();
                }
                [x, y, tile_id] => {
                    let tile_id: TileId = tile_id.try_into().expect("invalid tile id");
                    tiles.insert((x.clone(), y.clone()), tile_id);
                }
                _ => panic!("invalid chunk"),
            });
    }

    pub fn send_movement(&mut self, v: isize) {
        self.input.borrow_mut().push(v);
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
