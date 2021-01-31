use std::{thread, time::Duration};

use intcode_computer::Instruction;

use crate::simulation::{GameState, Simulation};
const DELAY: Duration = Duration::from_millis(20);

pub fn solve(mut program: Vec<isize>) -> isize {
    program[0] = 2;
    let mut simulation = Simulation::new(program);
    let mut brain = PlayerBrain::default();

    loop {
        let (state, instr) = simulation.execute();
        println!("{}", simulation);
        println!("Score: {}", simulation.score);
        if instr == Instruction::Halt {
            break;
        }
        let movement = brain.get_movement(state);

        simulation.send_movement(movement);
        simulation.computer.exec(&instr);
        thread::sleep(DELAY);
    }

    simulation.score
}

type Position = (isize, isize);
#[derive(Default)]
struct PlayerBrain {
    last_ball_position: Option<Position>,
    last_direction: Option<isize>,
}

impl PlayerBrain {
    fn get_movement(&mut self, s: GameState) -> isize {
        let next_ball_x = self.get_next_ball_x(&s.ball_pos);
        let x_delta = s.player_pos.0 - next_ball_x;
        self.last_ball_position = Some(s.ball_pos);
        // NOTE: I'm sorry, but here be dragons :D
        if next_ball_x == s.ball_pos.0 {
            return 0;
        }
        if (s.ball_pos.1 - s.player_pos.1).abs() > 1 && x_delta.abs() == 1 {
            return 0;
        }

        let direction = if x_delta < 0 {
            1
        } else if x_delta > 0 {
            -1
        } else {
            0
        };
        self.last_direction = Some(direction);

        direction
    }

    fn get_next_ball_x(&mut self, ball_pos: &Position) -> isize {
        let last_ball_x = self.last_ball_position.map(|p| p.0).unwrap_or(ball_pos.0);
        let x_delta = ball_pos.0 - last_ball_x;

        if x_delta == 0 {
            return match self.last_direction {
                Some(direction) => ball_pos.0 + direction * -1,
                None => ball_pos.0,
            };
        }

        ball_pos.0 + x_delta
    }
}
