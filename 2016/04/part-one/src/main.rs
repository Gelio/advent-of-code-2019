use std::collections::HashMap;
use std::io;
use std::cmp::Ordering;

fn main() {
    let mut line = String::new();
    let mut serial_id_sum = 0;

    loop {
        line.clear();
        match io::stdin().read_line(&mut line) {
            Ok(_) => (),
            Err(error) => {
                println!("Error: {}", error);
                std::process::exit(1);
            }
        }

        let line = line.trim();
        if line.len() == 0 {
            break;
        }

        let (hash, opening_bracket_index) = extract_hash(&line);


        let last_dash_index = line.rfind('-').unwrap();
        let serial_id =
            usize::from_str_radix(&line[last_dash_index + 1..opening_bracket_index], 10).unwrap();



        let room_name = &line[..last_dash_index].replace("-", "");
        let mut character_map: HashMap<_, _> = HashMap::new();



        for character in room_name.chars() {
            let mut previous_count = 0;
            if let Some(count) = character_map.get(&character) {
                previous_count = *count;
            }

            character_map.insert(character, previous_count + 1);
        }



        let mut character_map_vector: Vec<_> = character_map.iter().collect();
        character_map_vector.sort_by(|x, y| {
            let numbers_ordering = y.1.cmp(x.1);
            if numbers_ordering == Ordering::Equal {
                return x.0.cmp(y.0);
            }

            numbers_ordering
        });

        let correct_hash = &character_map_vector[..5];
        let mut correct = true;

        for (index, hash_character) in hash.chars().enumerate() {
            if hash_character != *correct_hash.get(index).unwrap().0 {
                correct = false;
                break;
            }
        }

        if correct {
            serial_id_sum += serial_id;
        }
    }

    println!("Serial ID sum: {}", serial_id_sum);
}

fn extract_hash(line: &str) -> (&str, usize) {
    let opening_bracket_index = line.find('[').unwrap();
    let closing_bracket_index = line.len() - 1;

    (
        &line[(opening_bracket_index + 1)..closing_bracket_index],
        opening_bracket_index,
    )
}
