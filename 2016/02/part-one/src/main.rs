use std::io;

#[derive(Copy, Clone)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

struct Position(usize, usize);

fn main() {
    let mut line = String::new();
    let numbers = [[1, 2, 3], [4, 5, 6], [7, 8, 9]];
    let mut position = Position(1, 1);
    let mut key_code: Vec<i32> = Vec::new();

    loop {
        line.clear();
        match io::stdin().read_line(&mut line) {
            Ok(_) => (),
            Err(error) => {
                println!("Input error: {}", error);
                std::process::exit(1);
            }
        }

        let line = line.trim();
        if line.len() == 0 {
            break;
        }

        for direction in line.chars().map(parse_direction) {
            move_finger(&mut position, direction);
        }

        let key_code_digit = numbers[position.0][position.1];
        key_code.push(key_code_digit);
    }

    println!("{:?}", key_code);
}

fn parse_direction(serialized_direction: char) -> Direction {
    match serialized_direction {
        'U' => Direction::Up,
        'D' => Direction::Down,
        'L' => Direction::Left,
        'R' => Direction::Right,
        _ => {
            println!("Invalid direction: {}", serialized_direction as i32);
            std::process::exit(1);
        }
    }
}

fn move_finger(position: &mut Position, direction: Direction) {
    match direction {
        Direction::Up => {
            if position.0 > 0 {
                position.0 -= 1;
            }
        },

        Direction::Down => {
            if position.0 < 2 {
                position.0 += 1;
            }
        },

        Direction::Right => {
            if position.1 < 2 {
                position.1 += 1;
            }
        }

        Direction::Left => {
            if position.1 > 0 {
                position.1 -= 1;
            }
        }
    }
}