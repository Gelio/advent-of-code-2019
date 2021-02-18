use std::io;

enum Direction {
    North,
    East,
    South,
    West,
}

enum TurnDirection {
    Right,
    Left
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

    let mut x = 0;
    let mut y = 0;
    let mut direction = Direction::North;

    for (turn_direction, distance) in buffer.split_whitespace().map(clean_up_input).map(parse_input_value) {
        direction = turn(direction, turn_direction);
        process_move(&mut x, &mut y, &direction, distance);
    }

    println!("x: {}, y: {}", x, y);
    let x_distance = x.abs();
    let y_distance = y.abs();
    let total_distance = x_distance + y_distance;
    
    println!("Total distance: {}", total_distance);
}

fn process_move(x: &mut i32, y: &mut i32, direction: &Direction, distance: i32) -> () {
    match *direction {
        Direction::North => *y += distance,
        Direction::East => *x += distance,
        Direction::South => *y -= distance,
        Direction::West => *x -= distance,
    }
}

fn parse_input_value(s: &str) -> (TurnDirection, i32) {
    let (serialized_direction, serialized_distance) = s.split_at(1);
    let turn_direction = match serialized_direction {
        "L" => TurnDirection::Left,
        "R" => TurnDirection::Right,
        _ => {
            println!("Turn direction parsing error in {}, unknown direction: {}", s, serialized_direction);
            std::process::exit(1);
        }
    };
    let distance: i32 = std::str::FromStr::from_str(serialized_distance).unwrap();

    (turn_direction, distance)
}

fn turn(initial_direction: Direction, turn_direction: TurnDirection) -> Direction {
    match initial_direction {
        Direction::North => {
            match turn_direction {
                TurnDirection::Left => Direction::West,
                TurnDirection::Right => Direction::East,
            }
        },
        Direction::East => {
            match turn_direction {
                TurnDirection::Left => Direction::North,
                TurnDirection::Right => Direction::South
            }
        },
        Direction::South => {
            match turn_direction {
                TurnDirection::Left => Direction::East,
                TurnDirection::Right => Direction::West,
            }
        },
        Direction::West => {
            match turn_direction {
                TurnDirection::Left => Direction::South,
                TurnDirection::Right => Direction::North
            }
        }
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
