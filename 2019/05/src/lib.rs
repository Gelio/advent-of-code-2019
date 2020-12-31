enum ArgMode {
    Position,
    Immediate,
}

impl ArgMode {
    fn parse(c: char) -> Self {
        match c {
            '0' => ArgMode::Position,
            '1' => ArgMode::Immediate,
            _ => panic!("Invalid arg mode {}", c),
        }
    }
}

#[derive(PartialEq, Debug)]
enum Instruction {
    Add {
        arg1: isize,
        arg2: isize,
        out: usize,
    },
    Multiply {
        arg1: isize,
        arg2: isize,
        out: usize,
    },
    ReadInput {
        to: usize,
    },
    WriteOutput {
        from: usize,
    },
    Halt,
}

pub struct Computer {
    memory: Vec<isize>,
    input: Box<dyn Iterator<Item = isize>>,
    pub output: Vec<isize>,
    ip: usize,
}

impl Computer {
    pub fn new(memory: Vec<isize>, input: Vec<isize>) -> Self {
        Self {
            memory,
            input: Box::new(input.into_iter()),
            output: Vec::new(),
            ip: 0,
        }
    }

    pub fn run(&mut self) {
        while self.step() {}
    }

    fn step(&mut self) -> bool {
        let instr = self.parse_instruction();
        match instr {
            Instruction::Add { arg1, arg2, out } => {
                self.memory[out] = arg1 + arg2;
                true
            }
            Instruction::Multiply { arg1, arg2, out } => {
                self.memory[out] = arg1 * arg2;
                true
            }
            Instruction::ReadInput { to } => {
                self.memory[to] = self.input.next().expect("Input too short");
                true
            }
            Instruction::WriteOutput { from } => {
                self.output.push(self.memory[from]);
                true
            }
            Instruction::Halt => false,
        }
    }

    fn parse_instruction(&mut self) -> Instruction {
        let instr = self.memory[self.ip];
        let instr_digits = format!("{:0>5}", instr.to_string());

        let opcode = instr_digits.get(3..).unwrap();

        match opcode {
            "01" => {
                let instr = Instruction::Add {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.memory[self.ip + 3] as usize,
                };

                self.ip += 4;

                instr
            }
            "02" => {
                let instr = Instruction::Multiply {
                    arg1: self.get_arg(&instr_digits, 1),
                    arg2: self.get_arg(&instr_digits, 2),
                    out: self.memory[self.ip + 3] as usize,
                };

                self.ip += 4;

                instr
            }
            "03" => {
                let instr = Instruction::ReadInput {
                    to: self.memory[self.ip + 1] as usize,
                };

                self.ip += 2;

                instr
            }
            "04" => {
                let instr = Instruction::WriteOutput {
                    from: self.memory[self.ip + 1] as usize,
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

    fn get_arg(&self, opcode: &str, arg_index: usize) -> isize {
        let arg_mode = opcode.as_bytes()[3 - arg_index] as char;
        let arg_mode = ArgMode::parse(arg_mode);
        let v = self.memory[self.ip + arg_index];

        match arg_mode {
            ArgMode::Immediate => v,
            ArgMode::Position => self.memory[v as usize],
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{Computer, Instruction};

    #[test]
    fn computer_correctly_parses_instructions() {
        let mut computer = Computer::new(vec![1002, 4, 3, 4, 33], vec![]);

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
    fn computer_correctly_invokes_instructions() {
        let mut computer = Computer::new(vec![1002, 4, 3, 4, 33], vec![]);

        assert!(computer.step());
        assert_eq!(computer.ip, 4, "Invalid IP");

        assert!(!computer.step());
    }
}
