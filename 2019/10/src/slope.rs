use crate::point::Point;

#[derive(Debug, Clone, Hash, PartialEq, Eq)]
pub struct Slope {
    x: i32,
    y: i32,
}

impl Slope {
    pub fn get(p1: &Point, p2: &Point) -> Self {
        let x_diff = p2.x - p1.x;
        let y_diff = p2.y - p1.y;

        let gcd = gcd(x_diff, y_diff).abs();

        Self {
            x: x_diff / gcd,
            y: y_diff / gcd,
        }
    }
}

fn gcd(mut a: i32, mut b: i32) -> i32 {
    while b != 0 {
        let tmp = a % b;
        a = b;
        b = tmp;
    }

    a
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

    #[test]
    fn points_on_same_line_but_different_sides_should_have_different_slopes() {
        let p1 = Point::new(1, 0);
        let p2 = Point::new(1, 2);
        let p3 = Point::new(1, -2);

        assert_ne!(Slope::get(&p1, &p2), Slope::get(&p1, &p3));
    }
}
