use std::collections::HashMap;

use crate::{move_emulator::Color, point::Point};

pub fn print_map(map: &HashMap<Point, Color>) {
    let (from, to) = get_bounds(map.keys());

    for y in (from.y..=to.y).rev() {
        for x in from.x..=to.x {
            let color = map.get(&Point { x, y }).cloned().unwrap_or_default();

            print!("{}", color.to_string());
        }

        println!("");
    }
}

fn get_bounds<'a>(points: impl Iterator<Item = &'a Point>) -> (Point, Point) {
    let mut from = Point::default();
    let mut to = Point::default();

    for p in points {
        if from.x > p.x {
            from.x = p.x;
        }

        if from.y > p.y {
            from.y = p.y;
        }

        if to.x < p.x {
            to.x = p.x;
        }

        if to.y < p.y {
            to.y = p.y;
        }
    }

    (from, to)
}
