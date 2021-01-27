use std::fmt::Display;

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

    pub fn add(p1: &Self, p2: &Self) -> Self {
        Self {
            x: p1.x + p2.x,
            y: p1.y + p2.y,
            z: p1.z + p2.z,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn correctly_parses_positions() {
        struct TestCase {
            input: &'static str,
            expected_result: Position,
        }
        let test_cases = vec![
            TestCase {
                input: "<x=14, y=15, z=-2>",
                expected_result: Position {
                    x: 14,
                    y: 15,
                    z: -2,
                },
            },
            TestCase {
                input: "<x=17, y=-3, z=4>",
                expected_result: Position { x: 17, y: -3, z: 4 },
            },
            TestCase {
                input: "<x=6, y=12, z=-13>",
                expected_result: Position {
                    x: 6,
                    y: 12,
                    z: -13,
                },
            },
            TestCase {
                input: "<x=-2, y=10, z=-8>",
                expected_result: Position {
                    x: -2,
                    y: 10,
                    z: -8,
                },
            },
        ];

        test_cases.into_iter().for_each(|test_case| {
            let parsed = Position::parse(test_case.input).unwrap();

            assert_eq!(test_case.expected_result, parsed);
        })
    }

    #[test]
    fn reports_errors_for_invalid_positions() {
        let invalid_inputs = vec![
            "not a valid input",
            "  <x=14, y=5, z=-2>",
            "<x=14, y=5, z=-2",
            "<x=14, y=5, z=->",
        ];

        invalid_inputs.into_iter().for_each(|input| {
            let res = Position::parse(input);

            assert!(res.is_err());
        })
    }

    #[test]
    fn adds_correctly() {
        let p1 = Position { x: 1, y: 2, z: 3 };
        let p2 = Position { x: 3, y: 2, z: 1 };

        let result = Position::add(&p1, &p2);
        assert_eq!(result, Position { x: 4, y: 4, z: 4 });
    }
}
