use crate::{
    direction::{Direction, Turn},
    point::Point,
};

pub struct Robot {
    pub position: Point,
    pub direction: Direction,
}

impl Default for Robot {
    fn default() -> Self {
        Self {
            position: Point::default(),
            direction: Direction::Up,
        }
    }
}

impl Robot {
    pub fn turn_and_forward(&mut self, turn: &Turn) {
        self.direction = self.direction.turn(turn);
        self.position = self.position.forward(self.direction);
    }
}
