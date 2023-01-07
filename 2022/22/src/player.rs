use crate::{map::Map, moves::TurnDirection};

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub enum Direction {
    North,
    South,
    East,
    West,
}

impl Direction {
    pub fn turn(self, direction: TurnDirection) -> Self {
        use Direction::*;

        match direction {
            TurnDirection::Clockwise => match self {
                North => East,
                East => South,
                South => West,
                West => North,
            },

            TurnDirection::Counterclockwise => match self {
                North => West,
                West => South,
                South => East,
                East => North,
            },
        }
    }
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub struct Position {
    pub x: usize,
    pub y: usize,
}

impl Position {
    pub fn initial_position_from_map(map: &Map) -> Self {
        let column_index = map
            .rows()
            .first()
            .and_then(|line| line.first_open_tile())
            .expect("First row did not contain any open tiles");

        Self {
            x: column_index,
            y: 0,
        }
    }
}

#[derive(Debug, PartialEq, Eq)]
pub struct Player {
    pub direction: Direction,
    pub position: Position,
}

impl Player {
    pub fn new(position: Position) -> Self {
        Self {
            position,
            direction: Direction::East,
        }
    }

    pub fn password(&self) -> usize {
        let row = self.position.y + 1;
        let column = self.position.x + 1;
        let facing = match self.direction {
            Direction::East => 0,
            Direction::South => 1,
            Direction::West => 2,
            Direction::North => 3,
        };

        1000 * row + 4 * column + facing
    }
}
