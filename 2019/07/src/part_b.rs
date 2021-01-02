use std::{cell::RefCell, rc::Rc};

use intcode_computer::{Computer, Instruction};

use crate::unique_number_sequence_generator::UniqueNumberSequenceGenerator;

const AMPLIFIERS_COUNT: usize = 5;

pub fn part_b(computer_memory: Vec<isize>) -> isize {
    let sequences = UniqueNumberSequenceGenerator::new((5, 9), AMPLIFIERS_COUNT);
    let mut max_output = isize::MIN;

    for seq in sequences {
        let inputs: Vec<Rc<RefCell<Vec<isize>>>> = seq
            .into_iter()
            .map(|s| Rc::new(RefCell::new(vec![s as isize])))
            .collect();

        inputs.last().unwrap().borrow_mut().push(0);
        let mut computers = (0..AMPLIFIERS_COUNT)
            .map(|i| {
                Computer::with_output(
                    computer_memory.clone(),
                    Rc::clone(&inputs[(AMPLIFIERS_COUNT + i - 1) % AMPLIFIERS_COUNT]),
                    Rc::clone(&inputs[i]),
                )
            })
            .collect::<Vec<_>>();

        let mut computer_i = 0;

        loop {
            let computer = computers.get_mut(computer_i).unwrap();

            let halted = loop {
                let instr = computer.parse_and_exec_once();
                match instr {
                    Instruction::Halt => break true,
                    Instruction::WriteOutput { .. } => {
                        computer_i = (computer_i + 1) % AMPLIFIERS_COUNT;
                        break false;
                    }
                    _ => {}
                };
            };

            if halted {
                break;
            }
        }

        let output = *computers.last().unwrap().output().last().unwrap();
        if output > max_output {
            max_output = output;
        }
    }

    max_output
}

#[cfg(test)]
mod tests {
    use super::part_b;
    #[test]
    fn correct_output() {
        assert_eq!(
            part_b(vec![
                3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28,
                -1, 28, 1005, 28, 6, 99, 0, 0, 5
            ]),
            139629729
        );

        assert_eq!(
            part_b(vec![
                3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001,
                54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53,
                55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10
            ]),
            18216
        );
    }
}
