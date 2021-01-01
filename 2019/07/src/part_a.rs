use std::{cell::RefCell, rc::Rc};

use aoc_2019_05::Computer;

const AMPLIFIERS_COUNT: usize = 5;

pub struct MaxSequenceFinder {
    computer_memory: Vec<isize>,
    max_signal: isize,
    phase_sequence: Vec<usize>,
    used_phase_settings: Vec<bool>,
}

impl MaxSequenceFinder {
    pub fn find_max_signal(computer_memory: Vec<isize>) -> isize {
        let mut finder = Self {
            computer_memory,
            max_signal: isize::MIN,
            phase_sequence: Vec::new(),
            used_phase_settings: vec![false; AMPLIFIERS_COUNT],
        };

        finder.try_phase_settings(0);

        finder.max_signal
    }

    fn try_phase_settings(&mut self, last_input_signal: isize) {
        for i in 0..AMPLIFIERS_COUNT {
            if self.used_phase_settings[i] {
                continue;
            }

            let input = vec![i as isize, last_input_signal];
            let mut computer =
                Computer::new(self.computer_memory.clone(), Rc::new(RefCell::new(input)));
            computer.run_till_halt();
            let output_signal = *computer.output.borrow().get(0).unwrap();

            self.phase_sequence.push(i);
            self.used_phase_settings[i] = true;

            if self.phase_sequence.len() == AMPLIFIERS_COUNT {
                if output_signal > self.max_signal {
                    self.max_signal = output_signal;
                }
            } else {
                self.try_phase_settings(output_signal);
            }

            self.phase_sequence.pop();
            self.used_phase_settings[i] = false;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::MaxSequenceFinder;

    #[test]
    fn gets_correct_max_thruster_signal() {
        assert_eq!(
            MaxSequenceFinder::find_max_signal(vec![
                3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0,
            ]),
            43210
        );
        assert_eq!(
            MaxSequenceFinder::find_max_signal(vec![
                3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4,
                23, 99, 0, 0,
            ]),
            54321
        );
        assert_eq!(
            MaxSequenceFinder::find_max_signal(vec![
                3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33,
                1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0
            ]),
            65210
        );
    }
}
