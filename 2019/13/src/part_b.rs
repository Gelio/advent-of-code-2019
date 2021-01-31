use std::io::{stdin, Read};

use crate::simulation::Simulation;

pub fn solve(mut program: Vec<isize>) -> isize {
    program[0] = 2;
    let mut simulation = Simulation::new(program);
    let mut stdin = stdin();

    let mut buf = [0; 1];
    loop {
        simulation.execute();
        println!("{}", simulation);
        println!("Score: {}", simulation.score);

        stdin.read_exact(&mut buf).unwrap();
        match buf[0] as char {
            'h' | 'a' => {
                simulation.send_movement(-1);
            }
            'l' | 'd' => {
                simulation.send_movement(1);
            }
            'q' => {
                break;
            }
            _ => {
                simulation.send_movement(0);
            }
        }
        buf[0] = 'n' as u8;
    }

    simulation.score
}
