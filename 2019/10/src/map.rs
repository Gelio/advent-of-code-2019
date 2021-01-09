use std::fmt::Display;

use crate::point::Point;

#[derive(Debug)]
pub struct AsteroidMap {
    pub asteroids: Vec<Point>,
}

#[derive(Debug, PartialEq)]
pub enum ParseError {
    UnknownCharacter(char),
}

impl Display for ParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::UnknownCharacter(c) => write!(f, "unknown character encountered {}", c),
        }
    }
}

const ASTEROID: char = '#';
const EMPTY_SPACE: char = '.';

impl AsteroidMap {
    pub fn parse(lines: &Vec<&str>) -> Result<Self, ParseError> {
        let mut asteroids: Vec<Point> = Vec::new();

        for (y, line) in lines.iter().enumerate() {
            for (x, c) in line.chars().enumerate() {
                match c {
                    ASTEROID => asteroids.push(Point::new(x as i32, y as i32)),
                    EMPTY_SPACE => (),
                    _ => return Err(ParseError::UnknownCharacter(c)),
                };
            }
        }

        Ok(Self { asteroids })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn parses_a_valid_map() {
        let lines: Vec<&str> = ".#..#
        .....
        #####
        ....#
        ...##"
            .split('\n')
            .map(|line| line.trim())
            .collect();

        let map = AsteroidMap::parse(&lines).expect("could not parse map");

        assert_eq!(map.asteroids.len(), 10, "invalid number of asteroids found");
        assert!(map.asteroids.contains(&Point::new(1, 0)));
        assert!(map.asteroids.contains(&Point::new(4, 0)));
        assert!(map.asteroids.contains(&Point::new(0, 2)));
        assert!(!map.asteroids.contains(&Point::new(1, 1)));
    }
}
