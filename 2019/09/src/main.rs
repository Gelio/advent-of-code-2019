use std::{cell::RefCell, fs::read_to_string, rc::Rc};

fn main() {
    let input = read_to_string("input.txt").unwrap();
    let program = input
        .trim()
        .split(",")
        .map(|x| x.parse().expect(&format!("Cannot parse {}", x)))
        .collect();
    let mut computer = aoc_2019_05::Computer::new(program, Rc::new(RefCell::new(vec![1])));
    computer.run_till_halt();

    println!("Result A: {:?}", computer.output.borrow());
}
