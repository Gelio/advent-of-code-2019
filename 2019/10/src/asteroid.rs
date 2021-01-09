use crate::map::AsteroidMap;
use crate::point::Point;
use crate::slope::Slope;

#[derive(Debug, PartialEq, Eq)]
pub struct VisibleAsteroid {
    pos: Point,
    distance: i32,
    slope: Slope,
}

pub struct Asteroid {
    pub pos: Point,
    pub visible_asteroids: Vec<VisibleAsteroid>,
}

pub fn get_asteroids_with_visible_neighbors(map: &AsteroidMap) -> Vec<Asteroid> {
    map.asteroids
        .iter()
        .map(|pos| {
            let mut visible_asteroids: Vec<VisibleAsteroid> = Vec::new();

            for other_asteroid_pos in map.asteroids.iter() {
                let slope = Slope::get(&pos, &other_asteroid_pos);
                let distance = Point::get_distance(&pos, other_asteroid_pos);

                match visible_asteroids.binary_search_by(|asteroid| asteroid.slope.cmp(&slope)) {
                    Ok(index) => {
                        let asteroid = visible_asteroids.get(index).unwrap();
                        if asteroid.distance > distance {
                            visible_asteroids[index] = VisibleAsteroid {
                                pos: other_asteroid_pos.clone(),
                                slope,
                                distance,
                            };
                        }
                    }
                    Err(possible_index) => visible_asteroids.insert(
                        possible_index,
                        VisibleAsteroid {
                            pos: other_asteroid_pos.clone(),
                            slope,
                            distance,
                        },
                    ),
                };
            }

            Asteroid {
                pos: pos.clone(),
                visible_asteroids,
            }
        })
        .collect()
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
