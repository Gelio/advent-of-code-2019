#[derive(PartialEq)]
pub enum ArgMode {
    Position,
    Immediate,
    Relative,
}

impl ArgMode {
    pub fn parse(c: char) -> Self {
        match c {
            '0' => ArgMode::Position,
            '1' => ArgMode::Immediate,
            '2' => ArgMode::Relative,
            _ => panic!("Invalid arg mode {}", c),
        }
    }
}

#[derive(PartialEq, Debug)]
pub enum Instruction {
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
        val: isize,
    },
    JumpIfTrue {
        arg: isize,
        destination: usize,
    },
    JumpIfFalse {
        arg: isize,
        destination: usize,
    },
    LessThan {
        arg1: isize,
        arg2: isize,
        out: usize,
    },
    Equals {
        arg1: isize,
        arg2: isize,
        out: usize,
    },
    AdjustRelativeBase {
        change: isize,
    },
    Halt,
}
