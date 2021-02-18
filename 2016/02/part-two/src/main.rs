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
    let characters = [['x', 'x', '1', 'x', 'x'], ['x', '2', '3', '4', 'x'], ['5', '6', '7', '8', '9'], ['x', 'A', 'B', 'C', 'x'], ['x', 'x', 'D', 'x', 'x']];
    let mut position = Position(2, 0);

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
            move_finger(&mut position, direction, &characters);
        } 

        let key_code_digit = characters[position.0][position.1];
        print!("{}", key_code_digit);
    }
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

fn move_finger(position: &mut Position, direction: Direction, characters: &[[char; 5]; 5]) {
    match direction {
        Direction::Up => {
            if position.0 > 0 && characters[position.0 - 1][position.1] != 'x' {
                position.0 -= 1;
            }
        },

        Direction::Down => {
            if position.0 < 4 && characters[position.0 + 1][position.1] != 'x' {
                position.0 += 1;
            }
        },

        Direction::Right => {
            if position.1 < 4 && characters[position.0][position.1 + 1] != 'x' {
                position.1 += 1;
            }
        }

        Direction::Left => {
            if position.1 > 0 && characters[position.0][position.1 - 1] != 'x' {
                position.1 -= 1;
            }
        }
    }
}