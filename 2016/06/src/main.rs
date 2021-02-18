use std::io;
use std::collections::HashMap;

fn main() {
    let mut line = String::new();
    let mut characters_frequency: Vec<HashMap<char, u32>> = Vec::new();

    loop {
        line.clear();

        io::stdin().read_line(&mut line)
            .expect("Error while reading line");
        
        let line = line.trim();
        let line_length = line.len();
        if line_length == 0 {
            break;
        }

        if characters_frequency.len() == 0 {
            characters_frequency.reserve(line_length);

            for _ in 0..line_length {
                characters_frequency.push(HashMap::new());
            }
        }

        for (position, character) in line.chars().enumerate() {
            let occurences = characters_frequency[position].entry(character).or_insert(0);
            *occurences += 1;
        }
    }

    let mut result = Vec::new();

    let is_part_one = false;

    for (position, occurrences_map) in characters_frequency.iter().enumerate() {
        let most_frequent_character_tuple;

        if is_part_one {
            most_frequent_character_tuple = occurrences_map.iter().max_by_key(|x| x.1).unwrap();
        } else {
            most_frequent_character_tuple = occurrences_map.iter().min_by_key(|x| x.1).unwrap();
        }
        result.push(*most_frequent_character_tuple.0);
        
        println!("The most frequent character at position {} is {} ({} occurrences)", position, most_frequent_character_tuple.0, most_frequent_character_tuple.1);
    }

    println!("Final answer: {:?}", result);
}
