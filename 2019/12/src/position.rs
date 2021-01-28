use std::fmt::Display;

use nom::{
    bytes::complete::{tag, take_while1},
    character::is_digit,
    combinator::opt,
    error::ErrorKind,
    sequence::preceded,
    IResult,
};
use regex::Regex;

#[derive(Clone, PartialEq, Eq, Debug, Default, Hash)]
pub struct Position {
    pub x: i32,
    pub y: i32,
    pub z: i32,
}

impl Display for Position {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({}, {}, {})", self.x, self.y, self.z)
    }
}

impl Position {
    pub fn parse(input: &str) -> Result<Self, String> {
        let position_regex = Regex::new(r"^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$")
            .expect("Could not compile position regexp");
        let captured_groups = position_regex
            .captures(input)
            .ok_or_else(|| format!("no matches found in {}", input))?;

        let x = captured_groups
            .get(1)
            .expect("Cannot unwrap x position")
            .as_str()
            .parse::<i32>()
            .map_err(|e| format!("cannot parse x in {}: {}", input, e.to_string()))?;

        let y = captured_groups
            .get(2)
            .expect("Cannot unwrap y position")
            .as_str()
            .parse::<i32>()
            .map_err(|e| format!("cannot parse y in {}: {}", input, e.to_string()))?;

        let z = captured_groups
            .get(3)
            .expect("Cannot unwrap z position")
            .as_str()
            .parse::<i32>()
            .map_err(|e| format!("cannot parse z in {}: {}", input, e.to_string()))?;

        Ok(Self { x, y, z })
    }

    #[allow(dead_code)]
    pub fn parse_nom(input: &str) -> Result<Self, String> {
        PositionParser::parse(input)
    }

    pub fn add(p1: &Self, p2: &Self) -> Self {
        Self {
            x: p1.x + p2.x,
            y: p1.y + p2.y,
            z: p1.z + p2.z,
        }
    }
}

// Parses using parser combinators (nom)
// https://github.com/Geal/nom
struct PositionParser;

impl PositionParser {
    fn parse(input: &str) -> Result<Position, String> {
        let input = input.as_bytes();

        match Self::parse_inner(input) {
            Ok((_, pos)) => Ok(pos),
            Err(e) => Err(e.to_string()),
        }
    }

    fn parse_inner(input: &[u8]) -> IResult<&[u8], Position> {
        // <x=14, y=15, z=-2>
        let (input, x) = preceded(tag("<x="), Self::integer)(input)?;
        let (input, y) = preceded(tag(", y="), Self::integer)(input)?;
        let (input, z) = preceded(tag(", z="), Self::integer)(input)?;
        let (input, _) = tag(">")(input)?;

        if !input.is_empty() {
            return Err(nom::Err::Failure(nom::error::Error::new(
                input,
                ErrorKind::NonEmpty,
            )));
        }

        Ok((input, Position { x, y, z }))
    }

    fn integer(input: &[u8]) -> IResult<&[u8], i32> {
        let (input, minus) = opt(nom::character::streaming::char('-'))(input)?;
        let sign_multiplier = minus.map_or(1, |_| -1);

        let (input, digits) = take_while1(is_digit)(input)?;
        let num = (std::str::from_utf8(digits).unwrap())
            .parse::<i32>()
            .map_err(|_| {
                nom::Err::Error(nom::error::Error::new(input, nom::error::ErrorKind::Digit))
            })?;

        Ok((input, sign_multiplier * num))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    struct PassingTestCase {
        input: &'static str,
        expected_result: Position,
    }

    fn get_passing_test_cases() -> Vec<PassingTestCase> {
        vec![
            PassingTestCase {
                input: "<x=14, y=15, z=-2>",
                expected_result: Position {
                    x: 14,
                    y: 15,
                    z: -2,
                },
            },
            PassingTestCase {
                input: "<x=17, y=-3, z=4>",
                expected_result: Position { x: 17, y: -3, z: 4 },
            },
            PassingTestCase {
                input: "<x=6, y=12, z=-13>",
                expected_result: Position {
                    x: 6,
                    y: 12,
                    z: -13,
                },
            },
            PassingTestCase {
                input: "<x=-2, y=10, z=-8>",
                expected_result: Position {
                    x: -2,
                    y: 10,
                    z: -8,
                },
            },
        ]
    }

    fn get_invalid_inputs() -> Vec<&'static str> {
        vec![
            "not a valid input",
            "  <x=14, y=5, z=-2>",
            "<x=14, y=5, z=-2",
            "<x=14, y=5, z=->",
        ]
    }

    mod regexp_parsing {
        use super::*;

        #[test]
        fn correctly_parses_positions() {
            get_passing_test_cases().into_iter().for_each(|test_case| {
                let parsed = Position::parse(test_case.input).unwrap();

                assert_eq!(test_case.expected_result, parsed);
            })
        }

        #[test]
        fn reports_errors_for_invalid_positions() {
            get_invalid_inputs().into_iter().for_each(|input| {
                let res = Position::parse(input);

                assert!(res.is_err());
            })
        }
    }
    mod parser_combinators {
        use super::*;

        #[test]
        fn correctly_parses_positions() {
            get_passing_test_cases().into_iter().for_each(|test_case| {
                let parsed = Position::parse_nom(test_case.input).unwrap();

                assert_eq!(test_case.expected_result, parsed);
            })
        }

        #[test]
        fn reports_errors_for_invalid_positions() {
            get_invalid_inputs().into_iter().for_each(|input| {
                let res = Position::parse_nom(input);

                assert!(res.is_err());
            })
        }
    }

    #[test]
    fn adds_correctly() {
        let p1 = Position { x: 1, y: 2, z: 3 };
        let p2 = Position { x: 3, y: 2, z: 1 };

        let result = Position::add(&p1, &p2);
        assert_eq!(result, Position { x: 4, y: 4, z: 4 });
    }
}
