#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub enum Move {
    Forward(usize),
    Turn(TurnDirection),
}

#[derive(Debug, PartialEq, Eq, Clone, Copy)]
pub enum TurnDirection {
    Clockwise,
    Counterclockwise,
}

pub mod parser {
    use nom::{
        branch::alt,
        character::{self, complete::char},
        combinator::map,
        multi::many1,
        IResult,
    };

    use super::*;

    fn turn_clockwise(input: &str) -> IResult<&str, TurnDirection> {
        map(char('R'), |_| TurnDirection::Clockwise)(input)
    }
    fn turn_counterclockwise(input: &str) -> IResult<&str, TurnDirection> {
        map(char('L'), |_| TurnDirection::Counterclockwise)(input)
    }
    fn turn_direction(input: &str) -> IResult<&str, TurnDirection> {
        alt((turn_clockwise, turn_counterclockwise))(input)
    }

    fn move_turn(input: &str) -> IResult<&str, Move> {
        map(turn_direction, Move::Turn)(input)
    }

    fn move_forward(input: &str) -> IResult<&str, Move> {
        map(character::complete::u32, |num| Move::Forward(num as usize))(input)
    }

    pub fn moves(input: &str) -> IResult<&str, Vec<Move>> {
        many1(alt((move_turn, move_forward)))(input)
    }

    #[cfg(test)]
    mod tests {
        use nom::combinator::all_consuming;

        use super::*;

        #[test]
        fn parses_moves_line() {
            let input = "10R5L5";

            assert_eq!(
                all_consuming(moves)(input),
                Ok((
                    "",
                    vec![
                        Move::Forward(10),
                        Move::Turn(TurnDirection::Clockwise),
                        Move::Forward(5),
                        Move::Turn(TurnDirection::Counterclockwise),
                        Move::Forward(5)
                    ]
                ))
            )
        }
    }
}
