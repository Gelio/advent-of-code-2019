use std::io;
use std::usize;

const MAP_SIZE: usize = 50000;

#[derive(PartialEq, Clone)]
enum Tile {
  None,
  FirstWire,
  SecondWire,
  Crossing,
}

enum Direction {
  Up,
  Down,
  Left,
  Right,
}

struct Instruction {
  direction: Direction,
  steps: usize,
}

fn parse_direction(direction: char) -> Option<Direction> {
  match direction {
    'U' => Some(Direction::Up),
    'D' => Some(Direction::Down),
    'L' => Some(Direction::Left),
    'R' => Some(Direction::Right),
    _ => None,
  }
}

impl Instruction {
  fn parse(instruction_string: &str) -> Option<Instruction> {
    if instruction_string.len() < 2 {
      return None;
    }

    let possible_direction = instruction_string.chars().nth(0).and_then(parse_direction);
    let direction = match possible_direction {
      Some(direction) => direction,
      None => return None,
    };

    let steps_str = &instruction_string[1..];
    let steps_parse_result = steps_str.parse::<usize>();

    let steps = match steps_parse_result {
      Ok(steps) => steps,
      Err(_) => return None,
    };

    Some(Instruction {
      direction: direction,
      steps: steps,
    })
  }
}

fn parse_instructions(line: &str) -> Option<Vec<Instruction>> {
  let instructions: Vec<Option<Instruction>> =
    line.trim().split(",").map(Instruction::parse).collect();

  // println!("{0} instructions", instructions.iter().count());

  let mut instructions_vec: Vec<Instruction> = Vec::new();

  for instruction in instructions {
    match instruction {
      Some(instruction) => instructions_vec.push(instruction),
      None => return None,
    };
  }

  // Some(instructions.iter().map(|&x| *x.as_ref().unwrap()).collect())
  Some(instructions_vec)
}

fn get_next_coordinates(x: usize, y: usize, direction: &Direction) -> (usize, usize) {
  match direction {
    Direction::Up => (x, y + 1),
    Direction::Down => (x, y - 1),
    Direction::Left => (x - 1, y),
    Direction::Right => (x + 1, y),
  }
}

fn main() {
  let mut first_line = String::new();
  io::stdin()
    .read_line(&mut first_line)
    .expect("Cannot first read line");
  let first_line_instructions = parse_instructions(&first_line);

  let mut second_line = String::new();
  io::stdin()
    .read_line(&mut second_line)
    .expect("Cannot second read line");
  let second_line_instructions = parse_instructions(&second_line);

  let mut map = construct_map(MAP_SIZE);

  match first_line_instructions {
    None => {
      println!("First line is not valid");
      return;
    }
    Some(instructions) => {
      let mut x = MAP_SIZE / 2;
      let mut y = MAP_SIZE / 2;
      map[x][y] = Tile::Crossing;

      for instruction in instructions {
        // println!("Got instruction");
        for _ in 0..instruction.steps {
          let (next_x, next_y) = get_next_coordinates(x, y, &instruction.direction);
          // println!("moving from {0} {1} to {2} {3}", x, y, next_x, next_y);
          x = next_x;
          y = next_y;
          map[x][y] = Tile::FirstWire;
        }
      }
    }
  }

  println!("Simulated first wire");

  match second_line_instructions {
    None => {
      println!("Second line is not valid");
      return;
    }
    Some(instructions) => {
      let mut x = MAP_SIZE / 2;
      let mut y = MAP_SIZE / 2;
      map[x][y] = Tile::Crossing;

      for instruction in instructions {
        for _ in 0..instruction.steps {
          let (next_x, next_y) = get_next_coordinates(x, y, &instruction.direction);
          x = next_x;
          y = next_y;

          map[x][y] = match map[x][y] {
            Tile::FirstWire => Tile::Crossing,
            _ => Tile::SecondWire,
          }
        }
      }
    }
  }

  println!("Simulated second wire");

  let center_x = MAP_SIZE / 2;
  let center_y = MAP_SIZE / 2;
  let (crossing_x, crossing_y, distance) = find_closest_crossing(&map, center_x, center_y);

  println!(
    "Closest crossing is at {0} {1} (distance: {2})",
    crossing_x, crossing_y, distance
  );
}

fn construct_map(map_size: usize) -> Vec<Vec<Tile>> {
  let mut map: Vec<Vec<Tile>> = Vec::new();
  let mut row: Vec<Tile> = Vec::new();
  for _ in 0..map_size {
    row.push(Tile::None);
  }

  for _ in 1..map_size {
    map.push(row.clone());
  }

  map
}

fn find_closest_crossing(
  map: &Vec<Vec<Tile>>,
  target_x: usize,
  target_y: usize,
) -> (usize, usize, usize) {
  let mut crossing_x = 0;
  let mut crossing_y = 0;
  let mut distance = usize::MAX;

  let mut try_position = |x: usize, y: usize| {
    if map[x][y] != Tile::Crossing {
      return false;
    }

    let current_distance = get_distance(x, y, target_x, target_y);
    if current_distance < distance {
      crossing_x = x;
      crossing_y = y;
      distance = current_distance;

      return true;
    }

    false
  };

  for distance in 1..map.len() / 2 {
    // Top-right diagonal
    let mut x = target_x;
    let mut y = target_y - distance;

    while y <= target_y {
      if try_position(x, y) {
        return (x, y, distance);
      }
      y += 1;
      x += 1;
    }

    // Bottom-right diagonal
    let mut x = target_x + distance;
    let mut y = target_y;

    while x >= target_x {
      if try_position(x, y) {
        return (x, y, distance);
      }
      y += 1;
      x -= 1;
    }

    // Bottom-left diagonal
    let mut x = target_x;
    let mut y = target_y + distance;

    while y >= target_y {
      if try_position(x, y) {
        return (x, y, distance);
      }
      y -= 1;
      x -= 1;
    }

    // Bottom-left diagonal
    let mut x = target_x - distance;
    let mut y = target_y;

    while x <= target_x {
      if try_position(x, y) {
        return (x, y, distance);
      }
      y -= 1;
      x += 1;
    }
  }

  // for x in 0..map.len() {
  //   for y in 0..map[x].len() {
  //     if map[x][y] != Tile::Crossing || (x == target_x && y == target_y) {
  //       continue;
  //     }

  //     let current_distance = get_distance(x, y, target_x, target_y);
  //     if current_distance < distance {
  //       crossing_x = x;
  //       crossing_y = y;
  //       distance = current_distance;
  //     }
  //   }
  // }

  (crossing_x, crossing_y, distance)
}

fn get_distance(x1: usize, y1: usize, x2: usize, y2: usize) -> usize {
  ((x1 as i32 - x2 as i32).abs() + (y1 as i32 - y2 as i32).abs()) as usize
}
