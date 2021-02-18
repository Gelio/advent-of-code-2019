#[derive(Debug)]
pub enum ChipTransferTarget {
    Output(usize),
    Bot(usize),
}

#[derive(Debug)]
pub struct ChipTransferInstruction {
    pub low_value_target: ChipTransferTarget,
    pub high_value_target: ChipTransferTarget,
}

#[derive(Debug)]
pub struct Bot {
    id: usize,
    high_value_chip: Option<i32>,
    low_value_chip: Option<i32>,
    transfer_instructions: Vec<ChipTransferInstruction>,
}

impl Bot {
    pub fn new(id: usize) -> Bot {
        Bot {
            id,
            low_value_chip: None,
            high_value_chip: None,
            transfer_instructions: Vec::new(),
        }
    }

    pub fn equip_chip(&mut self, chip_value: i32) {
        if let Some(low_value_chip) = self.low_value_chip {
            if let Some(high_value_chip) = self.high_value_chip {
                panic!(
                    "Bot {} received a chip {} but it already holds two chips {} {}",
                    self.id,
                    chip_value,
                    low_value_chip,
                    high_value_chip
                );
            }

            if chip_value > low_value_chip {
                self.high_value_chip = Some(chip_value);
            } else {
                self.high_value_chip = self.low_value_chip;
                self.low_value_chip = Some(chip_value);
            }
        } else if let Some(high_value_chip) = self.high_value_chip {
            if chip_value < high_value_chip {
                self.low_value_chip = Some(chip_value);
            } else {
                self.low_value_chip = self.high_value_chip;
                self.high_value_chip = Some(chip_value);
            }
        } else {
            self.low_value_chip = Some(chip_value);
        }

        // println!("Bot equipped chips: {:?}", self);
    }

    pub fn get_low_value_chip(&self) -> Option<i32> {
        self.low_value_chip
    }

    pub fn get_high_value_chip(&self) -> Option<i32> {
        self.high_value_chip
    }

    pub fn has_both_chips(&self) -> bool {
        self.low_value_chip.is_some() && self.high_value_chip.is_some()
    }

    pub fn has_instructions(&self) -> bool {
        !self.transfer_instructions.is_empty()
    }

    pub fn can_execute_instruction(&self) -> bool {
        self.has_both_chips() && self.has_instructions()
    }

    pub fn clear(&mut self) {
        self.low_value_chip = None;
        self.high_value_chip = None;
    }

    pub fn add_new_instruction(&mut self, instruction: ChipTransferInstruction) {
        self.transfer_instructions.push(instruction);
    }

    pub fn pop_instruction(&mut self) -> ChipTransferInstruction {
        self.transfer_instructions.remove(0)
    }
}
