use std::collections;
use std::io;

use collections::HashMap;

enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn parse(direction: char) -> Option<Self> {
        match direction {
            'U' => Some(Direction::Up),
            'D' => Some(Direction::Down),
            'L' => Some(Direction::Left),
            'R' => Some(Direction::Right),
            _ => None,
        }
    }
}

struct Instruction {
    direction: Direction,
    steps: usize,
}

impl Instruction {
    fn parse(instruction_string: &str) -> Option<Instruction> {
        if instruction_string.len() < 2 {
            return None;
        }

        let direction = instruction_string
            .chars()
            .nth(0)
            .and_then(Direction::parse)?;

        let steps_str = &instruction_string[1..];
        let steps = steps_str.parse::<usize>().ok()?;

        Some(Instruction { direction, steps })
    }
}

fn parse_instructions(line: &str) -> Option<Vec<Instruction>> {
    line.trim().split(",").map(Instruction::parse).collect()
}

fn get_next_coordinates(x: i64, y: i64, direction: &Direction) -> (i64, i64) {
    match direction {
        Direction::Up => (x, y + 1),
        Direction::Down => (x, y - 1),
        Direction::Left => (x - 1, y),
        Direction::Right => (x + 1, y),
    }
}

#[derive(Debug)]
struct WireCrossing {
    x: i64,
    y: i64,
    wires_distance: usize,
}

fn main() {
    let first_line_instructions = {
        let mut first_line = String::new();
        io::stdin()
            .read_line(&mut first_line)
            .expect("Cannot first read line");

        parse_instructions(&first_line)
    }
    .unwrap_or(Vec::new());

    let second_line_instructions = {
        let mut second_line = String::new();
        io::stdin()
            .read_line(&mut second_line)
            .expect("Cannot second read line");

        parse_instructions(&second_line)
    }
    .unwrap_or(Vec::new());

    let mut first_wire_positions: HashMap<(i64, i64), usize> = HashMap::new();
    first_line_instructions
        .iter()
        .fold((0, 0, 0), |(x, y, dist), instruction| {
            let mut curr_x = x;
            let mut curr_y = y;
            let mut curr_dist = dist;

            for _ in 0..instruction.steps {
                let next_pos = get_next_coordinates(curr_x, curr_y, &instruction.direction);
                curr_x = next_pos.0;
                curr_y = next_pos.1;
                curr_dist += 1;
                first_wire_positions.insert((curr_x, curr_y), curr_dist);
            }

            (curr_x, curr_y, curr_dist)
        });

    let mut second_wire_positions: HashMap<(i64, i64), usize> = HashMap::new();
    let mut crossings: Vec<WireCrossing> = Vec::new();

    second_line_instructions
        .iter()
        .fold((0, 0, 0), |(x, y, dist), instruction| {
            let mut curr_x = x;
            let mut curr_y = y;
            let mut curr_dist = dist;

            for _ in 0..instruction.steps {
                let next_pos = get_next_coordinates(curr_x, curr_y, &instruction.direction);
                curr_x = next_pos.0;
                curr_y = next_pos.1;
                curr_dist += 1;
                second_wire_positions.insert((curr_x, curr_y), curr_dist);

                match first_wire_positions.get(&(curr_x, curr_y)) {
                    Some(&first_wire_dist) => crossings.push(WireCrossing {
                        x: curr_x,
                        y: curr_y,
                        wires_distance: curr_dist + first_wire_dist,
                    }),

                    None => (),
                }
            }

            (curr_x, curr_y, curr_dist)
        });

    let closest_crossing =
        crossings
            .iter()
            .skip(1)
            .fold(crossings.get(0), |closest_opt, crossing| {
                closest_opt.map(|closest| {
                    if closest.wires_distance < crossing.wires_distance {
                        closest
                    } else {
                        crossing
                    }
                })
            });

    match closest_crossing {
        Some(crossing) => {
            println!("Closest crossing: {:?}", crossing)
        }
        None => {
            println!("Crossing not found")
        }
    }
}
