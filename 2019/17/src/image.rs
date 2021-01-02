#[derive(PartialEq, Debug)]
pub struct Point(pub usize, pub usize);

impl std::fmt::Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}

#[derive(PartialEq, Debug)]
pub enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(PartialEq, Debug)]
pub struct Robot {
    pos: Point,
    direction: Direction,
}

#[derive(PartialEq, Debug)]
pub enum Tile {
    Empty,
    Scaffold,
}

#[allow(non_snake_case)]
pub fn parse_1D_image(image: &Vec<isize>) -> (Vec<Vec<Tile>>, Option<Robot>) {
    let mut x: usize = 0;
    let mut y: usize = 0;

    let mut image_2d: Vec<Vec<Tile>> = vec![vec![]];
    let mut robot: Option<Robot> = None;

    image
        .iter()
        .map(|c| *c as u8 as char)
        .for_each(|c| match c {
            '#' => {
                image_2d[y].push(Tile::Scaffold);
                x += 1;
            }
            '.' => {
                image_2d[y].push(Tile::Empty);
                x += 1;
            }
            '\n' => {
                image_2d.push(Vec::new());
                y += 1;
                x = 0;
            }
            '^' => {
                image_2d[y].push(Tile::Scaffold);
                robot = Some(Robot {
                    pos: Point(x, y),
                    direction: Direction::Up,
                });
                x += 1;
            }
            'V' => {
                image_2d[y].push(Tile::Scaffold);
                robot = Some(Robot {
                    pos: Point(x, y),
                    direction: Direction::Down,
                });
                x += 1;
            }
            '<' => {
                image_2d[y].push(Tile::Scaffold);
                robot = Some(Robot {
                    pos: Point(x, y),
                    direction: Direction::Left,
                });
                x += 1;
            }
            '>' => {
                image_2d[y].push(Tile::Scaffold);
                robot = Some(Robot {
                    pos: Point(x, y),
                    direction: Direction::Right,
                });
                x += 1;
            }
            _ => panic!("Unknown tile {} at pos {:?}", c, Point(x, y)),
        });

    (image_2d, robot)
}

pub fn find_intersections(image: &Vec<Vec<Tile>>) -> Vec<Point> {
    let height = image.len();
    let width = image[0].len();
    let mut neighboring_scaffolds: Vec<Vec<i32>> = vec![vec![0; width]; height];

    image.iter().enumerate().for_each(|(y, row)| {
        row.iter()
            .enumerate()
            .filter(|(_, t)| **t == Tile::Scaffold)
            .for_each(|(x, _)| {
                if x > 0 {
                    neighboring_scaffolds[y][x - 1] += 1;
                }
                if y > 0 {
                    neighboring_scaffolds[y - 1][x] += 1;
                }
                if x + 1 < width {
                    neighboring_scaffolds[y][x + 1] += 1;
                }
                if y + 1 < height {
                    neighboring_scaffolds[y + 1][x] += 1;
                }
            });
    });

    let mut intersections: Vec<Point> = Vec::new();
    neighboring_scaffolds
        .iter()
        .enumerate()
        .for_each(|(y, row)| {
            row.iter().enumerate().for_each(|(x, &neighbors)| {
                if neighbors == 4 && image[y][x] == Tile::Scaffold {
                    intersections.push(Point(x, y));
                }
            })
        });

    intersections
}

#[cfg(test)]
mod tests {
    use std::iter;

    use super::*;

    fn get_image_from_string(str: &str) -> Vec<isize> {
        str.split("\n")
            .flat_map(|line| {
                line.trim()
                    .as_bytes()
                    .iter()
                    .map(|&c| c as isize)
                    .chain(iter::once('\n' as u8 as isize))
            })
            .collect()
    }

    #[test]
    fn finds_robot() {
        let image = get_image_from_string(
            "..#..........
            ..#..........
            #######...###
            #.#...#...#.#
            #############
            ..#...#...#..
            ..#####...^..",
        );

        let (_, robot) = parse_1D_image(&image);

        let expected_robot = Robot {
            pos: Point(10, 6),
            direction: Direction::Up,
        };

        assert_eq!(robot, Some(expected_robot), "invalid robot");
    }
    #[test]
    fn finds_intersections() {
        let image = get_image_from_string(
            "..#..........
            ..#..........
            #######...###
            #.#...#...#.#
            #############
            ..#...#...#..
            ..#####...^..",
        );
        let (image, _) = parse_1D_image(&image);
        let intersections = find_intersections(&image);

        assert_eq!(intersections.len(), 4, "invalid number of intersections");
        assert!(intersections.contains(&Point(2, 2)));
        assert!(intersections.contains(&Point(2, 4)));
        assert!(intersections.contains(&Point(6, 4)));
        assert!(intersections.contains(&Point(10, 4)));
    }
}
