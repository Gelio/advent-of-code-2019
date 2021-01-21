use std::{
    cell::RefCell,
    collections::HashMap,
    convert::{TryFrom, TryInto},
    rc::Rc,
};

use intcode_computer::{Computer, Instruction};

use crate::{direction::Turn, point::Point, robot::Robot};

#[derive(Debug, PartialEq, Eq, Copy, Clone)]
pub enum Color {
    Black = 0,
    White = 1,
}

impl TryFrom<isize> for Color {
    type Error = String;

    fn try_from(value: isize) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Color::Black),
            1 => Ok(Color::White),
            x => Err(format!("invalid color: {}", x)),
        }
    }
}

impl Default for Color {
    fn default() -> Self {
        Self::Black
    }
}

impl Into<isize> for &Color {
    fn into(self) -> isize {
        *self as isize
    }
}

pub struct MoveEmulator {
    pub map: HashMap<Point, Color>,
    pub robot: Robot,
    pub computer: Computer,
    input: Rc<RefCell<Vec<isize>>>,
}

impl MoveEmulator {
    pub fn new(program: Vec<isize>) -> Self {
        let input = Rc::default();
        let computer = Computer::new(program, Rc::clone(&input));

        Self {
            map: HashMap::new(),
            robot: Robot::default(),
            input,
            computer,
        }
    }

    pub fn run_till_halt(&mut self) {
        loop {
            let instr = self.perform_single_move();

            if let Instruction::Halt = instr {
                break;
            }
        }
    }

    fn perform_single_move(&mut self) -> Instruction {
        let color_under_robot = self
            .map
            .get(&self.robot.position)
            .unwrap_or(&Color::default())
            .clone();

        self.input.borrow_mut().push((&color_under_robot).into());
        let mut turn: Option<Turn> = None;
        let mut color_to_paint: Option<Color> = None;
        let instr = loop {
            let instr = self.computer.parse_and_exec_once();
            match instr {
                Instruction::WriteOutput { val } => {
                    if color_to_paint.is_none() {
                        color_to_paint = Some(val.try_into().expect("invalid color to paint"));
                    } else {
                        turn = Some(val.try_into().expect("invalid turn"));
                        break instr;
                    }
                }
                Instruction::Halt => break instr,
                _ => (),
            }
        };

        match (turn, color_to_paint) {
            (Some(turn), Some(color_to_paint)) => {
                self.map.insert(self.robot.position, color_to_paint);
                self.robot.turn_and_forward(&turn);
            }
            _ => {}
        }

        instr
    }
}
