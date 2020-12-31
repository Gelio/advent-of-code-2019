pub struct UniqueNumberSequenceGenerator {
    phase_sequence: Vec<usize>,
    used_phase_settings: Vec<bool>,
    phase_settings_range: (usize, usize),
    possible_phase_settings: Vec<Vec<usize>>,
    sequence_len: usize,
}

impl Iterator for UniqueNumberSequenceGenerator {
    type Item = Vec<usize>;

    fn next(&mut self) -> Option<Self::Item> {
        if self.get_next_phase_setting() {
            return Some(self.phase_sequence.clone());
        }

        None
    }
}

impl UniqueNumberSequenceGenerator {
    pub fn new(range: (usize, usize), sequence_len: usize) -> Self {
        Self {
            phase_sequence: Vec::new(),
            phase_settings_range: range,
            used_phase_settings: vec![false; range.1 + 1],
            sequence_len,
            possible_phase_settings: vec![(range.0..=range.1).rev().collect()],
        }
    }

    fn get_next_phase_setting(&mut self) -> bool {
        loop {
            self.phase_sequence.pop().and_then(|s| {
                self.used_phase_settings[s] = false;
                Some(s)
            });

            loop {
                let possible_phase_settings = match self.possible_phase_settings.last_mut() {
                    Some(x) => x,
                    // No possible settings
                    None => return false,
                };

                let phase_setting = match possible_phase_settings.pop() {
                    Some(x) => x,
                    None => {
                        // No possibilies left for the current amplifier
                        self.possible_phase_settings.pop();
                        break;
                    }
                };

                self.phase_sequence.push(phase_setting);
                self.used_phase_settings[phase_setting] = true;

                if self.phase_sequence.len() != self.sequence_len {
                    let possible_next_phase_settings = (self.phase_settings_range.0
                        ..=self.phase_settings_range.1)
                        .rev()
                        .filter(|x| !self.used_phase_settings[*x])
                        .collect();

                    self.possible_phase_settings
                        .push(possible_next_phase_settings);

                    continue;
                }

                return true;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::UniqueNumberSequenceGenerator;

    #[test]
    fn generates_valid_sequences() {
        assert_eq!(
            UniqueNumberSequenceGenerator::new((0, 2), 2).collect::<Vec<Vec<usize>>>(),
            vec![
                vec![0, 1],
                vec![0, 2],
                vec![1, 0],
                vec![1, 2],
                vec![2, 0],
                vec![2, 1],
            ]
        );
    }
}
