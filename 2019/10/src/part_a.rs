use crate::asteroid::{get_asteroids_with_visible_neighbors, Asteroid};
use crate::map::AsteroidMap;

pub fn get_best_asteroid_neighbors_count(map: &AsteroidMap) -> usize {
    get_best_asteroid(map)
        .expect("no asteroids found")
        .visible_asteroids
        .len()
}

fn get_best_asteroid(map: &AsteroidMap) -> Option<Asteroid> {
    let asteroids = get_asteroids_with_visible_neighbors(&map);

    asteroids
        .into_iter()
        .max_by(|a1, a2| a1.visible_asteroids.len().cmp(&a2.visible_asteroids.len()))
}

#[cfg(test)]
mod tests {
    use crate::point::Point;

    use super::*;

    fn parse_map(input: &str) -> AsteroidMap {
        AsteroidMap::parse(&input.split_whitespace().collect()).expect("could not parse map")
    }

    struct TestCase {
        input: &'static str,
        expected_visible_asteroids: usize,
        best_asteroid_pos: Point,
    }

    #[test]
    fn best_asteroid_examples() {
        let cases = vec![
            TestCase {
                input: ".#..#
                .....
                #####
                ....#
                ...##",
                expected_visible_asteroids: 8,
                best_asteroid_pos: Point::new(3, 4),
            },
            TestCase {
                input: "......#.#.
            #..#.#....
            ..#######.
            .#.#.###..
            .#..#.....
            ..#....#.#
            #..#....#.
            .##.#..###
            ##...#..#.
            .#....####",
                expected_visible_asteroids: 33,
                best_asteroid_pos: Point::new(5, 8),
            },
            TestCase {
                input: "#.#...#.#.
                .###....#.
                .#....#...
                ##.#.#.#.#
                ....#.#.#.
                .##..###.#
                ..#...##..
                ..##....##
                ......#...
                .####.###.",
                expected_visible_asteroids: 35,
                best_asteroid_pos: Point::new(1, 2),
            },
            TestCase {
                input: ".#..#..###
                ####.###.#
                ....###.#.
                ..###.##.#
                ##.##.#.#.
                ....###..#
                ..#.#..#.#
                #..#.#.###
                .##...##.#
                .....#.#..",
                expected_visible_asteroids: 41,
                best_asteroid_pos: Point::new(6, 3),
            },
            TestCase {
                input: ".#..##.###...#######
                ##.############..##.
                .#.######.########.#
                .###.#######.####.#.
                #####.##.#.##.###.##
                ..#####..#.#########
                ####################
                #.####....###.#.#.##
                ##.#################
                #####.##.###..####..
                ..######..##.#######
                ####.##.####...##..#
                .#####..#.######.###
                ##...#.##########...
                #.##########.#######
                .####.#.###.###.#.##
                ....##.##.###..#####
                .#.#.###########.###
                #.#.#.#####.####.###
                ###.##.####.##.#..##",
                expected_visible_asteroids: 210,
                best_asteroid_pos: Point::new(11, 13),
            },
        ];

        for (i, case) in cases.iter().enumerate() {
            let map = parse_map(case.input);
            let asteroid = get_best_asteroid(&map).expect(&format!("{}: no asteroids found", i));

            assert_eq!(
                asteroid.visible_asteroids.len(),
                case.expected_visible_asteroids,
                "{}: invalid number of visible asteroids",
                i
            );
            assert_eq!(
                asteroid.pos, case.best_asteroid_pos,
                "{}: invalid position",
                i
            );
        }
    }
}
