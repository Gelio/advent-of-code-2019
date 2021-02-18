pub fn supports_ssl(address: &str) -> bool {
  let (supernet_sequence, hypernet_sequence) = split_address_sequences(address);
  let mut supernet_sequence_part = &supernet_sequence[..];

  loop {
    match contains_aba(&supernet_sequence_part) {
      None => break,
      Some((aba, match_index)) => {
        if contains_bab(&hypernet_sequence, aba) {
          return true;
        }

        supernet_sequence_part = &supernet_sequence_part[match_index + 1..];
      }
    }
  }


  false
}

fn split_address_sequences(address: &str) -> (String, String) {
  let mut supernet_sequence = String::new();
  let mut hypernet_sequence = String::new();
  let mut address_part = address;

  loop {
    let opening_square_bracket_match = address_part.find('[');
    if let None = opening_square_bracket_match {
      break;
    }

    let opening_square_bracket_index = opening_square_bracket_match.unwrap();

    supernet_sequence.push_str(&address_part[..opening_square_bracket_index]);
    let closing_square_bracket_index = address_part.find(']').unwrap();
    
    hypernet_sequence.push_str(&address_part[opening_square_bracket_index + 1 .. closing_square_bracket_index]);
    address_part = &address_part[closing_square_bracket_index + 1..];
  }

  supernet_sequence.push_str(address_part);
  (supernet_sequence, hypernet_sequence)
}

fn contains_aba(address: &str) -> Option<(&str, usize)> {
    let characters: Vec<char> = address.chars().collect();

    if address.len() < 3 {
      return None;
    }

    for index in 0..address.len() - 2 {
        if characters[index] == characters[index + 2] &&
            characters[index] != characters[index + 1]
        {
            return Some((&address[index..index + 3], index));
        }
    }

    None
}

fn contains_bab(address: &str, aba: &str) -> bool {
    let address_characters: Vec<char> = address.chars().collect();
    let aba_characters: Vec<char> = aba.chars().collect();

    for index in 0..address.len() - 2 {
        if address_characters[index] == aba_characters[1] &&
            address_characters[index + 1] == aba_characters[0] &&
            address_characters[index + 2] == aba_characters[1]
        {
            return true;
        }
    }

    false
}