use crate::{asteroid::get_visible_asteroids, map::AsteroidMap, point::Point};

// Removes asteroids visible from the laser position and returns the removed ones
fn remove_visible_asteroids(map: &mut AsteroidMap, laser_pos: &Point) -> Vec<Point> {
    let visible_asteroids = get_visible_asteroids(map, laser_pos);

    let mut aligned_visible_asteroids: Vec<_> = visible_asteroids
        .iter()
        .map(|(slope, asteroid)| (slope.angle(), asteroid))
        .collect();

    aligned_visible_asteroids.sort_by(|(s1, _), (s2, _)| (*s1).partial_cmp(s2).unwrap());

    let first_asteroid_to_remove_index = aligned_visible_asteroids
        .iter()
        // The "up" direction is at 270 degrees in the coordinate system.
        // Coordinate system's X direction is right, Y direction is down.
        .position(|(slope, _)| *slope >= 270.0)
        .unwrap_or(0);

    aligned_visible_asteroids.rotate_left(first_asteroid_to_remove_index);

    for (_, asteroid) in aligned_visible_asteroids.iter() {
        let index = map
            .asteroids
            .iter()
            .position(|pos| pos.eq(&asteroid.pos))
            .unwrap();

        map.asteroids.swap_remove(index);
    }

    aligned_visible_asteroids
        .iter()
        .map(|(_, asteroid)| asteroid.pos.clone())
        .collect()
}

const TARGET_ASTEROID_TO_REMOVE: usize = 200;
pub fn find_200th_removed_asteroid(map: &mut AsteroidMap, laser_pos: &Point) -> Point {
    let mut removed_asteroids: Vec<Point> = Vec::new();

    while removed_asteroids.len() < TARGET_ASTEROID_TO_REMOVE {
        removed_asteroids.append(&mut remove_visible_asteroids(map, laser_pos))
    }

    removed_asteroids.swap_remove(TARGET_ASTEROID_TO_REMOVE - 1)
}

#[cfg(test)]
mod tests {
    use std::vec;

    use crate::part_a::get_best_asteroid;

    use super::*;

    #[test]
    fn removes_visible_asteroids() {
        let mut map = AsteroidMap::parse(
            &".#....#####...#..
            ##...##.#####..##
            ##...#...#.#####.
            ..#.....#...###..
            ..#.#.....#....##"
                .split_whitespace()
                .collect(),
        )
        .expect("cannot parse map");

        let laser_pos = Point::new(8, 3);

        let removed_asteroids = remove_visible_asteroids(&mut map, &laser_pos);

        let expected_removed_asteroids_prefix = vec![
            Point::new(8, 1),
            Point::new(9, 0),
            Point::new(9, 1),
            Point::new(10, 0),
            Point::new(9, 2),
        ];

        assert_eq!(
            removed_asteroids[0..expected_removed_asteroids_prefix.len()],
            expected_removed_asteroids_prefix[..]
        );
    }

    #[test]
    fn correct_200th_asteroid() {
        let mut map = AsteroidMap::parse(
            &".#..##.###...#######
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
            ###.##.####.##.#..##"
                .split_whitespace()
                .collect(),
        )
        .expect("cannot parse map");

        let laser_pos = get_best_asteroid(&map)
            .expect("cannot find laser position")
            .pos;

        let target_asteroid = find_200th_removed_asteroid(&mut map, &laser_pos);

        assert_eq!(target_asteroid, Point::new(8, 2))
    }
}
