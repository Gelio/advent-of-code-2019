use std::{cell::RefCell, fs::read_to_string, mem, rc::Rc};

use aoc_2019_05::{Computer, Instruction};

fn main() {
    let input = read_to_string("input.txt").unwrap();

    let memory: Vec<isize> = input
        .trim()
        .split(",")
        .map(|x| x.parse().expect(&format!("Cannot parse number {}", x)))
        .collect();

    let mut computer = Computer::new(memory.clone(), Rc::new(RefCell::new(vec![1])));
    computer.run_till(mem::discriminant(&Instruction::Halt));

    println!("Result A: {:?}", computer.output.borrow());

    let mut computer = Computer::new(memory, Rc::new(RefCell::new(vec![5])));
    computer.run_till(mem::discriminant(&Instruction::Halt));

    println!("Result B: {:?}", computer.output.borrow());
}
