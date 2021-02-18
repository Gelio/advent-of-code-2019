extern crate regex;

use std::io;

mod factory;
use factory::Factory;



fn main() {
    let mut factory = Factory::new();
    let mut line = String::new();

    loop {
        line.clear();
        io::stdin().read_line(&mut line).expect("Cannot read line");
        let line = line.trim();
        if line.len() == 0 {
            break;
        }

        factory.interpret_instruction(line);
    }
    println!("Read all lines");

    let outputs = &factory.outputs;
    println!("Result from 0, 1 and 2 outputs: {}", outputs[0][0] * outputs[1][0] * outputs[2][0]);
}
