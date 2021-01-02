use std::{cell::RefCell, rc::Rc};

use crate::instruction::{ArgMode, Instruction};
use crate::ram::RAM;

pub struct Computer {
    ram: RAM,
    input: Rc<RefCell<Vec<isize>>>,
    input_index: usize,
    relative_base: isize,
    output: Rc<RefCell<Vec<isize>>>,
    ip: usize,
}

impl Computer {
    pub fn new(memory: Vec<isize>, input: Rc<RefCell<Vec<isize>>>) -> Self {
        Self::with_output(memory, input, Rc::new(RefCell::new(Vec::new())))
    }

    pub fn with_output(
        program: Vec<isize>,
        input: Rc<RefCell<Vec<isize>>>,
        output: Rc<RefCell<Vec<isize>>>,
    ) -> Self {
        Self {
            ram: RAM::new(program),
            input,
            output,
            input_index: 0,
            relative_base: 0,
            ip: 0,
        }
    }

    pub fn with_empty_input(memory: Vec<isize>) -> Self {
        Self::new(memory, Rc::new(RefCell::new(Vec::new())))
    }

    pub fn run_till_halt(&mut self) {
        while self.parse_and_exec_once() != Instruction::Halt {}
    }

    pub fn parse_and_exec_once(&mut self) -> Instruction {
        let instr = self.parse_instruction();
        self.exec(&instr);

        instr
    }

    pub fn output(&self) -> std::cell::Ref<Vec<isize>> {
        self.output.borrow()
    }

    fn exec(&mut self, instr: &Instruction) {
        match *instr {
            Instruction::Add { arg1, arg2, out } => {
                self.ram.set(out, arg1 + arg2);
            }
            Instruction::Multiply { arg1, arg2, out } => {
                self.ram.set(out, arg1 * arg2);
            }
            Instruction::ReadInput { to } => {
                let value_read = *self
                    .input
                    .borrow()
                    .get(self.input_index)
                    .expect("Input too short");
                self.ram.set(to, value_read);
                self.input_index += 1;
            }
            Instruction::WriteOutput { val } => {
                self.output.borrow_mut().push(val);
            }
            Instruction::JumpIfTrue { arg, destination } => {
                if arg != 0 {
                    self.ip = destination;
                }
            }
            Instruction::JumpIfFalse { arg, destination } => {
                if arg == 0 {
                    self.ip = destination;
                }
            }
            Instruction::LessThan { arg1, arg2, out } => {
                let result = if arg1 < arg2 { 1 } else { 0 };
                self.ram.set(out, result);
            }
            Instruction::Equals { arg1, arg2, out } => {
                let result = if arg1 == arg2 { 1 } else { 0 };
                self.ram.set(out, result);
            }
            Instruction::AdjustRelativeBase { change } => {
                self.relative_base += change;
            }
            Instruction::Halt => {}
        }
    }

    fn parse_instruction(&mut self) -> Instruction {
        let instr = self.ram.get(self.ip);
        let instr_digits = format!("{:0>5}", instr.to_string());

        let opcode = instr_digits.get(3..).unwrap();

        match opcode {
            "01" => {
                let instr = Instruction::Add {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.get_addr_arg(&instr_digits, 3),
                };

                self.ip += 4;

                instr
            }
            "02" => {
                let instr = Instruction::Multiply {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.get_addr_arg(&instr_digits, 3),
                };

                self.ip += 4;

                instr
            }
            "03" => {
                let instr = Instruction::ReadInput {
                    to: self.get_addr_arg(&instr_digits, 1),
                };

                self.ip += 2;

                instr
            }
            "04" => {
                let instr = Instruction::WriteOutput {
                    val: self.get_arg(&instr_digits, 1),
                };

                self.ip += 2;

                instr
            }
            "05" => {
                let instr = Instruction::JumpIfTrue {
                    arg: self.get_arg(&instr_digits, 1),
                    destination: self.get_arg(&instr_digits, 2) as usize,
                };

                self.ip += 3;

                instr
            }
            "06" => {
                let instr = Instruction::JumpIfFalse {
                    arg: self.get_arg(&instr_digits, 1),
                    destination: self.get_arg(&instr_digits, 2) as usize,
                };

                self.ip += 3;

                instr
            }
            "07" => {
                let instr = Instruction::LessThan {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.get_addr_arg(&instr_digits, 3),
                };

                self.ip += 4;

                instr
            }
            "08" => {
                let instr = Instruction::Equals {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.get_addr_arg(&instr_digits, 3),
                };

                self.ip += 4;

                instr
            }
            "09" => {
                let instr = Instruction::AdjustRelativeBase {
                    change: self.get_arg(&instr_digits, 1),
                };

                self.ip += 2;

                instr
            }
            "99" => {
                self.ip += 1;

                Instruction::Halt
            }
            _ => panic!("Invalid opcode {}", opcode),
        }
    }

    fn get_arg(&mut self, instr_digits: &str, arg_index: usize) -> isize {
        let arg_mode = instr_digits.as_bytes()[3 - arg_index] as char;
        let arg_mode = ArgMode::parse(arg_mode);

        self.get_arg_with_mode(arg_mode, arg_index)
    }

    fn get_arg_with_mode(&mut self, mode: ArgMode, arg_index: usize) -> isize {
        let v = self.ram.get(self.ip + arg_index);

        match mode {
            ArgMode::Immediate => v,
            ArgMode::Position => self.ram.get(v as usize),
            ArgMode::Relative => self.ram.get((v + self.relative_base) as usize),
        }
    }

    fn get_addr_arg(&mut self, instr_digits: &str, arg_index: usize) -> usize {
        let arg_mode = instr_digits.as_bytes()[3 - arg_index] as char;
        let arg_mode = ArgMode::parse(arg_mode);

        let mut addr = self.get_arg_with_mode(ArgMode::Immediate, arg_index);
        if arg_mode == ArgMode::Relative {
            addr += self.relative_base;
        }

        addr as usize
    }
}

#[cfg(test)]
mod tests {
    use std::{cell::RefCell, rc::Rc};

    use crate::instruction::Instruction;

    use super::Computer;

    #[test]
    fn correctly_parses_basic_multiply_instruction() {
        let mut computer = Computer::with_empty_input(vec![1002, 4, 3, 4, 33]);

        let instr = computer.parse_instruction();
        assert_eq!(
            instr,
            Instruction::Multiply {
                arg1: 33,
                arg2: 3,
                out: 4
            }
        );
    }

    #[test]
    fn correctly_invokes_basic_instructions() {
        let mut computer = Computer::with_empty_input(vec![1002, 4, 3, 4, 33]);

        computer.parse_and_exec_once();
        assert_eq!(computer.ip, 4, "Invalid IP");

        assert_eq!(computer.parse_and_exec_once(), Instruction::Halt);
    }

    #[test]
    fn jump_position_mode() {
        let input = vec![3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![0])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![0], "jump 0 test");

        let mut computer = Computer::new(input, Rc::new(RefCell::new(vec![5])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![1], "jump non-0 test");
    }

    #[test]
    fn jump_immediate_mode() {
        let input = vec![3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![0])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![0], "jump 0 test");

        let mut computer = Computer::new(input, Rc::new(RefCell::new(vec![5])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![1], "jump non-0 test");
    }

    #[test]
    fn part_b_larger_example() {
        let input = vec![
            3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0,
            0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4,
            20, 1105, 1, 46, 98, 99,
        ];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![7])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![999], "below 8");

        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![8])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![1000], "8");

        let mut computer = Computer::new(input, Rc::new(RefCell::new(vec![9])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), vec![1001], "above 9");
    }

    #[test]
    fn relative_base_clone_input() {
        let input = vec![
            109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99,
        ];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![])));
        computer.run_till_halt();
        assert_eq!(computer.output().clone(), input);
    }

    #[test]
    fn output_16_digit_number() {
        let input = vec![1102, 34915192, 34915192, 7, 4, 7, 99, 0];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![])));
        computer.run_till_halt();
        let output = computer.output();

        assert_eq!(output.len(), 1, "invalid length");

        let output_num = output[0].to_string();
        assert_eq!(output_num.len(), 16, "invalid output number");
    }

    #[test]
    fn handle_large_numbers() {
        let input = vec![104, 1125899906842624, 99];
        let mut computer = Computer::new(input.clone(), Rc::new(RefCell::new(vec![])));
        computer.run_till_halt();
        let output = computer.output.borrow();

        assert_eq!(output.len(), 1, "invalid length");

        let output_num = output[0];
        assert_eq!(output_num, 1125899906842624, "invalid output number");
    }
}
