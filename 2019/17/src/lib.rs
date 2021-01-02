pub mod image;

use image::{find_intersections, parse_1D_image, Point};
use intcode_computer::Computer;

pub fn part_1(program: Vec<isize>) -> usize {
    let mut computer = Computer::with_empty_input(program);
    computer.run_till_halt();

    let (image, _) = parse_1D_image(&computer.output());
    let intersections = find_intersections(&image);

    intersections.iter().map(|&Point(x, y)| x * y).sum()
}
