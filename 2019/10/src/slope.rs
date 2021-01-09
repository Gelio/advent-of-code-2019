use std::{cmp::Ordering, mem};

use crate::point::Point;

#[derive(Debug, PartialOrd)]
pub enum Slope {
    Up,
    Down,
    Num(f64),
}

impl PartialEq for Slope {
    fn eq(&self, other: &Self) -> bool {
        if mem::discriminant(self) != mem::discriminant(other) {
            return false;
        }

        if let Self::Num(x) = self {
            if let Self::Num(y) = other {
                return floats_equal(*x, *y);
            }
        }

        return true;
    }
}

impl Eq for Slope {}

impl Ord for Slope {
    fn cmp(&self, other: &Self) -> Ordering {
        self.partial_cmp(other).unwrap()
    }
}

const EPS: f64 = 0.00001;
fn floats_equal(x: f64, y: f64) -> bool {
    (x - y).abs() < EPS
}

impl Slope {
    pub fn get(p1: &Point, p2: &Point) -> Self {
        let x_diff = p2.x - p1.x;
        let y_diff = p2.y - p1.y;
        if x_diff == 0 {
            if y_diff > 0 {
                return Self::Down;
            } else {
                return Self::Up;
            }
        }

        Self::Num(y_diff as f64 / x_diff as f64)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn points_on_the_same_line_should_have_same_slope() {
        let slopes = vec![(1, 2), (-5, 0), (0, 5), (27, 11)];

        for (x_delta, y_delta) in slopes {
            let base = Point::new(1, 3);
            let expected_slope = Slope::get(&base, &Point::new(base.x + x_delta, base.y + y_delta));

            for i in 1..=100 {
                let p = Point::new(base.x + i * x_delta, base.y + i * y_delta);
                let slope = Slope::get(&base, &p);

                assert_eq!(
                    slope, expected_slope,
                    "incorrect slope for i={} when slope was ({}, {})",
                    i, x_delta, y_delta
                );
            }

            let expected_slope = Slope::get(&base, &Point::new(base.x - x_delta, base.y - y_delta));

            for i in -100..0 {
                let p = Point::new(base.x + i * x_delta, base.y + i * y_delta);
                let slope = Slope::get(&base, &p);

                assert_eq!(
                    slope, expected_slope,
                    "incorrect slope for i={} when slope was ({}, {})",
                    i, x_delta, y_delta
                );
            }
        }
    }

    #[test]
    fn points_on_vertical_line_should_have_same_slope() {
        let p1 = Point::new(1, 0);
        let p2 = Point::new(1, 2);
        let p3 = Point::new(1, 4);

        assert_eq!(Slope::get(&p1, &p2), Slope::get(&p1, &p3));
    }
}
