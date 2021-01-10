use std::collections::HashMap;

use crate::map::AsteroidMap;
use crate::point::Point;
use crate::slope::Slope;

#[derive(Debug, PartialEq, Eq)]
pub struct VisibleAsteroid {
    pub pos: Point,
    distance: i32,
}

#[derive(Debug, PartialEq, Eq)]
pub struct Asteroid {
    pub pos: Point,
    pub visible_asteroids: HashMap<Slope, VisibleAsteroid>,
}

pub fn get_asteroids_with_visible_neighbors(map: &AsteroidMap) -> Vec<Asteroid> {
    map.asteroids
        .iter()
        .map(|pos| Asteroid {
            pos: pos.clone(),
            visible_asteroids: get_visible_asteroids(map, pos),
        })
        .collect()
}

pub fn get_visible_asteroids(map: &AsteroidMap, pos: &Point) -> HashMap<Slope, VisibleAsteroid> {
    let mut visible_asteroids: HashMap<Slope, VisibleAsteroid> = HashMap::new();

    for other_asteroid_pos in map.asteroids.iter() {
        if pos == other_asteroid_pos {
            continue;
        }

        let slope = Slope::get(pos, other_asteroid_pos);
        let distance = Point::get_distance(pos, other_asteroid_pos);

        match visible_asteroids.get(&slope) {
            Some(asteroid) if asteroid.distance < distance => {}
            _ => {
                visible_asteroids.insert(
                    slope,
                    VisibleAsteroid {
                        pos: other_asteroid_pos.clone(),
                        distance,
                    },
                );
            }
        };
    }

    visible_asteroids
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn returns_correct_number_of_visible_asteroids() {
        let lines: Vec<&str> = ".#..#
        .....
        #####
        ....#
        ...##"
            .split_whitespace()
            .collect();
        let map = AsteroidMap::parse(&lines).expect("cannot parse map");

        let asteroids = get_asteroids_with_visible_neighbors(&map);

        let asteroid = asteroids
            .iter()
            .find(|a| a.pos == Point::new(3, 4))
            .expect("asteroid not found");

        assert_eq!(
            asteroid.visible_asteroids.len(),
            8,
            "invalid number of visible asteroids"
        );
    }
}
