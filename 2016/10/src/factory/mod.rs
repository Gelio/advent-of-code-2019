use regex::Regex;

mod bot;
use self::bot::*;

pub struct Factory {
  bots: Vec<Bot>,
  pub outputs: Vec<Vec<i32>>,
}

impl Factory {
  pub fn new() -> Factory {
    let mut bots: Vec<Bot> = Vec::new();
    for id in 0..220 {
      bots.push(Bot::new(id));
    }

    let mut outputs: Vec<Vec<i32>> = Vec::new();
    for _ in 0..220 {
      outputs.push(Vec::new());
    }

    Factory { bots, outputs }
  }

  pub fn interpret_instruction(&mut self, input: &str) {
    let assign_chip_regex = Regex::new(r"value (\d+) goes to bot (\d+)").unwrap();
    let move_chips_regex = Regex::new(
      r"bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)",
    ).unwrap();
    let mut bot_id: usize = 0;

    for capture in assign_chip_regex.captures_iter(input) {
      let chip_value: i32 = capture[1].parse().unwrap();
      bot_id = capture[2].parse().unwrap();
      self.bots[bot_id].equip_chip(chip_value);
    }

    for capture in move_chips_regex.captures_iter(input) {
      bot_id = capture[1].parse().unwrap();
      let low_chip_transfer_target = self.parse_chip_transfer_target(&capture[2], &capture[3]);
      let high_chip_transfer_target = self.parse_chip_transfer_target(&capture[4], &capture[5]);
      let transfer_instruction = ChipTransferInstruction {
        low_value_target: low_chip_transfer_target,
        high_value_target: high_chip_transfer_target,
      };

      let bot = &mut self.bots[bot_id];
      bot.add_new_instruction(transfer_instruction);
    }

    if self.bots[bot_id].can_execute_instruction() {
      self.execute_bot_instruction(bot_id);
    }
  }

  fn parse_chip_transfer_target(&mut self, target: &str, target_id: &str) -> ChipTransferTarget {
    let target_id = target_id.parse().unwrap();

    if target.eq("output") {
      ChipTransferTarget::Output(target_id)
    } else {
      ChipTransferTarget::Bot(target_id)
    }
  }

  fn execute_bot_instruction(&mut self, bot_id: usize) {
    let (instruction, low_value_chip, high_value_chip) = {
      let bot = &mut self.bots[bot_id];
      let instruction = bot.pop_instruction();
      let low_value_chip = bot.get_low_value_chip().unwrap();
      let high_value_chip = bot.get_high_value_chip().unwrap();
      bot.clear();
      (instruction, low_value_chip, high_value_chip)
    };

    if low_value_chip == 17 && high_value_chip == 61 {
      println!("Bot {} is responsible for handling of 17 and 61 chips", bot_id);
    }

    match instruction.low_value_target {
      ChipTransferTarget::Bot(other_bot_id) => {
        self.bots[other_bot_id].equip_chip(low_value_chip);
      },
      ChipTransferTarget::Output(output_id) => {
        self.outputs[output_id].push(low_value_chip);
      }
    }

    match instruction.high_value_target {
      ChipTransferTarget::Bot(other_bot_id) => {
        self.bots[other_bot_id].equip_chip(high_value_chip);
      },
      ChipTransferTarget::Output(output_id) => {
        self.outputs[output_id].push(high_value_chip);
      }
    }

    if let ChipTransferTarget::Bot(other_bot_id) = instruction.low_value_target {
      if self.bots[other_bot_id].can_execute_instruction() {
        self.execute_bot_instruction(other_bot_id);
      }
    }

    if let ChipTransferTarget::Bot(other_bot_id) = instruction.high_value_target {
      if self.bots[other_bot_id].can_execute_instruction() {
        self.execute_bot_instruction(other_bot_id);
      }
    }    
  }
}
