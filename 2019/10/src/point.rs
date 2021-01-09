#[derive(Debug, PartialEq, Eq, Clone)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl From<(i32, i32)> for Point {
    fn from((x, y): (i32, i32)) -> Self {
        Self { x, y }
    }
}

impl Point {
    pub fn new(x: i32, y: i32) -> Self {
        Self { x, y }
    }

    pub fn get_distance(p1: &Point, p2: &Point) -> i32 {
        (p2.x - p1.x).abs() + (p2.y - p1.y).abs()
    }
}
