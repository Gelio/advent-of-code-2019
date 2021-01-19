use std::convert::{TryFrom, TryInto};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Direction {
    Up = 0,
    Right = 1,
    Down = 2,
    Left = 3,
}

impl TryFrom<isize> for Direction {
    type Error = String;

    fn try_from(value: isize) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Self::Up),
            1 => Ok(Self::Right),
            2 => Ok(Self::Down),
            3 => Ok(Self::Left),
            x => Err(format!("invalid direction: {}", x)),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Turn {
    Left,
    Right,
}

impl TryFrom<isize> for Turn {
    type Error = String;

    fn try_from(value: isize) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Self::Left),
            1 => Ok(Self::Right),
            x => Err(format!("invalid turn: {}", x)),
        }
    }
}

impl Direction {
    pub fn turn(&self, turn: &Turn) -> Self {
        let value_to_add = if *turn == Turn::Left { -1 } else { 1 };
        let b = *self as isize;
        let modulo = Direction::Left as isize + 1;

        ((b + value_to_add + modulo) % (modulo)).try_into().unwrap()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn correct_turns() {
        assert_eq!(Direction::Up.turn(&Turn::Right), Direction::Right);
        assert_eq!(Direction::Up.turn(&Turn::Left), Direction::Left);
        assert_eq!(Direction::Left.turn(&Turn::Left), Direction::Down);
        assert_eq!(Direction::Left.turn(&Turn::Right), Direction::Up);
    }
}
