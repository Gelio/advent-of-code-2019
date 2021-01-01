pub struct UniqueNumberSequenceGenerator {
    num_sequence: Vec<usize>,
    used_numbers: Vec<bool>,
    numbers_range: (usize, usize),
    possible_numbers_for_places: Vec<Vec<usize>>,
    sequence_len: usize,
}

impl Iterator for UniqueNumberSequenceGenerator {
    type Item = Vec<usize>;

    fn next(&mut self) -> Option<Self::Item> {
        if self.get_next_sequence() {
            return Some(self.num_sequence.clone());
        }

        None
    }
}

impl UniqueNumberSequenceGenerator {
    pub fn new(range: (usize, usize), sequence_len: usize) -> Self {
        Self {
            num_sequence: Vec::new(),
            numbers_range: range,
            used_numbers: vec![false; range.1 + 1],
            sequence_len,
            possible_numbers_for_places: vec![(range.0..=range.1).rev().collect()],
        }
    }

    fn get_next_sequence(&mut self) -> bool {
        loop {
            self.num_sequence.pop().and_then(|s| {
                self.used_numbers[s] = false;
                Some(s)
            });

            loop {
                let possible_numbers = match self.possible_numbers_for_places.last_mut() {
                    Some(x) => x,
                    // No possible numbers in the whole sequence
                    None => return false,
                };

                let num = match possible_numbers.pop() {
                    Some(x) => x,
                    None => {
                        // No possibilies left in the current place
                        self.possible_numbers_for_places.pop();
                        break;
                    }
                };

                self.num_sequence.push(num);
                self.used_numbers[num] = true;

                if self.num_sequence.len() != self.sequence_len {
                    let possible_numbers_for_next_place = (self.numbers_range.0
                        ..=self.numbers_range.1)
                        .rev()
                        .filter(|x| !self.used_numbers[*x])
                        .collect();

                    self.possible_numbers_for_places
                        .push(possible_numbers_for_next_place);

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
