pub fn supports_tls(address: &str) -> bool {
    !is_hypernet_address(address) && contains_abba(address)
}

fn is_hypernet_address(address: &str) -> bool {
    let mut address_part = address;

    loop {
        let opening_bracket_match = address_part.find('[');
        if let None = opening_bracket_match {
            break;
        }

        let opening_bracket_index = opening_bracket_match.unwrap();
        address_part = &address_part[opening_bracket_index + 1..];

        let closing_bracket_index = address_part.find(']').unwrap();

        if contains_abba(&address_part[..closing_bracket_index]) {
            return true;
        }
    }

    false
}

fn contains_abba(address: &str) -> bool {
    let characters: Vec<char> = address.chars().collect();

    for index in 0..address.len() - 3 {
        if characters[index] == characters[index + 3] &&
            characters[index + 1] == characters[index + 2] &&
            characters[index] != characters[index + 1]
        {
            return true;
        }
    }

    false
}