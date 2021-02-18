use std::io;
use std::collections::HashSet;

enum Direction {
    North,
    East,
    South,
    West,
}

enum TurnDirection {
    Right,
    Left,
}

#[derive(Hash, Eq, PartialEq, Debug)]
struct Position {
    x: i32,
    y: i32,
}

fn main() {
    let mut buffer = String::new();
    match io::stdin().read_line(&mut buffer) {
        Ok(_) => (),
        Err(err) => {
            println!("error: {:?}", err);
            std::process::exit(1)
        }
    }

    let mut position = Position { x: 0, y: 0 };
    let mut direction = Direction::North;
    let mut visited_positions: HashSet<Position> = HashSet::new();


    for (turn_direction, distance) in buffer
        .split_whitespace()
        .map(clean_up_input)
        .map(parse_input_value)
    {
        println!("{:?}", position);
        let mut distance = distance;
        direction = turn(direction, turn_direction);

        while distance > 0 {
            if visited_positions.contains(&position) {
                println!("Position visited twice: {:?}", position);
                break;
            }

            let new_position = process_move(&position, &direction, 1);
            visited_positions.insert(position);
            position = new_position;
            distance = distance - 1;
        }
    }

    if visited_positions.contains(&position) {
        println!("Position visited twice: {:?}", position);
    }

    println!("x: {}, y: {}", position.x, position.y);
    let x_distance = position.x.abs();
    let y_distance = position.y.abs();
    let total_distance = x_distance + y_distance;

    println!("Total distance: {}", total_distance);
}

fn process_move(position: &Position, direction: &Direction, distance: i32) -> Position {
    match *direction {
        Direction::North => Position {
            x: position.x,
            y: position.y + distance,
        },
        Direction::East => Position {
            x: position.x + distance,
            y: position.y,
        },
        Direction::South => Position {
            x: position.x,
            y: position.y - distance,
        },
        Direction::West => Position {
            x: position.x - distance,
            y: position.y,
        },
    }
}

fn parse_input_value(s: &str) -> (TurnDirection, i32) {
    let (serialized_direction, serialized_distance) = s.split_at(1);
    let turn_direction = match serialized_direction {
        "L" => TurnDirection::Left,
        "R" => TurnDirection::Right,
        _ => {
            println!(
                "Turn direction parsing error in {}, unknown direction: {}",
                s,
                serialized_direction
            );
            std::process::exit(1);
        }
    };
    let distance: i32 = std::str::FromStr::from_str(serialized_distance).unwrap();

    (turn_direction, distance)
}

fn turn(initial_direction: Direction, turn_direction: TurnDirection) -> Direction {
    match initial_direction {
        Direction::North => match turn_direction {
            TurnDirection::Left => Direction::West,
            TurnDirection::Right => Direction::East,
        },
        Direction::East => match turn_direction {
            TurnDirection::Left => Direction::North,
            TurnDirection::Right => Direction::South,
        },
        Direction::South => match turn_direction {
            TurnDirection::Left => Direction::East,
            TurnDirection::Right => Direction::West,
        },
        Direction::West => match turn_direction {
            TurnDirection::Left => Direction::South,
            TurnDirection::Right => Direction::North,
        },
    }
}

fn clean_up_input(s: &str) -> &str {
    if let Some(index) = s.find(",") {
        if let Some(cleaned_up_s) = s.get(..index) {
            return cleaned_up_s;
        }
    }

    s
}
