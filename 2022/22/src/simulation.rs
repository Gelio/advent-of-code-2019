use crate::{
    map::Map,
    moves::Move,
    player::{Direction, Player, Position},
};

#[derive(Debug)]
pub struct Simulation {
    player: Player,
    map: Map,
}

impl Simulation {
    pub fn new(map: Map) -> Self {
        Self {
            player: Player::new(Position::initial_position_from_map(&map)),
            map,
        }
    }

    pub fn simulate_move(&mut self, m: Move) {
        match m {
            Move::Turn(turn_direction) => {
                self.player.direction = self.player.direction.turn(turn_direction);
            }
            Move::Forward(steps) => {
                match self.player.direction {
                    Direction::North => {
                        // column negative
                        self.player.position.y = self.map.columns()[self.player.position.x]
                            .move_negative(self.player.position.y, steps);
                    }
                    Direction::East => {
                        // row positive
                        self.player.position.x = self.map.rows()[self.player.position.y]
                            .move_positive(self.player.position.x, steps);
                    }
                    Direction::South => {
                        // column positive
                        self.player.position.y = self.map.columns()[self.player.position.x]
                            .move_positive(self.player.position.y, steps);
                    }
                    Direction::West => {
                        // row negative
                        self.player.position.x = self.map.rows()[self.player.position.y]
                            .move_negative(self.player.position.x, steps);
                    }
                }
            }
        }
    }

    pub fn player(&self) -> &Player {
        &self.player
    }
}

#[cfg(test)]
const EXAMPLE_INPUT: &str = "        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5";

pub mod parser {
    use nom::{
        character::complete::newline,
        combinator::map,
        sequence::{preceded, tuple},
        IResult,
    };

    use crate::{map_tiles::parser::map_tiles, moves::parser::moves};

    use super::*;

    pub fn map_with_moves(input: &str) -> IResult<&str, (Map, Vec<Move>)> {
        tuple((
            map(map_tiles, Map::new),
            preceded(tuple((newline, newline)), moves),
        ))(input)
    }

    #[cfg(test)]
    mod tests {
        use nom::combinator::all_consuming;

        use super::*;

        #[test]
        fn parses_example() {
            all_consuming(map_with_moves)(EXAMPLE_INPUT).expect("Parsing error");
        }
    }
}

#[cfg(test)]
mod tests {
    use nom::combinator::all_consuming;

    use super::{parser::map_with_moves, *};

    #[test]
    fn example() {
        let (_, (map, moves)) =
            all_consuming(map_with_moves)(EXAMPLE_INPUT).expect("Parsing error");

        let mut simulation = Simulation::new(map);
        assert_eq!(simulation.player.position, Position { x: 8, y: 0 });

        for m in moves {
            simulation.simulate_move(m);
        }

        assert_eq!(
            simulation.player,
            Player {
                direction: Direction::East,
                position: Position { x: 7, y: 5 }
            }
        );

        assert_eq!(simulation.player.password(), 6032);
    }
}
