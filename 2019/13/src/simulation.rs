use std::{
    cell::RefCell,
    collections::HashMap,
    convert::{TryFrom, TryInto},
    fmt::Display,
    rc::Rc,
};

use intcode_computer::{Computer, Instruction};

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

pub struct GameState {
    pub player_pos: (isize, isize),
    pub ball_pos: (isize, isize),
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

    pub fn execute(&mut self) -> (GameState, Instruction) {
        let instr = loop {
            let instr = self.computer.parse_instruction();
            match instr {
                intcode_computer::Instruction::ReadInput { .. }
                | intcode_computer::Instruction::Halt => {
                    break instr;
                }
                _ => {
                    self.computer.exec(&instr);
                }
            }
        };

        let mut player_pos: Option<(isize, isize)> = None;
        let mut ball_pos: Option<(isize, isize)> = None;
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
                    match tile_id {
                        TileId::Ball => ball_pos = Some((x.clone(), y.clone())),
                        TileId::HorizontalPaddle => player_pos = Some((x.clone(), y.clone())),
                        _ => {}
                    }
                    tiles.insert((x.clone(), y.clone()), tile_id);
                }
                _ => panic!("invalid chunk"),
            });

        (
            GameState {
                player_pos: player_pos.expect("player not found"),
                ball_pos: ball_pos.expect("ball not found"),
            },
            instr,
        )
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
