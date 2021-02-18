use std::io;

fn main() {
    let mut line = String::new();

    io::stdin()
        .read_line(&mut line)
        .expect("Error while reading line");

    println!("Result's length: {}", get_decompressed_length(line.trim()));
}

fn get_decompressed_length(input: &str) -> u64 {
    let input_bytes = input.as_bytes();
    let mut current_index: usize = 0;
    let mut length = 0;

    while current_index < input.len() {
        let current_character = input_bytes[current_index] as char;

        if current_character == '(' {
            let separator_index = input[current_index..].find('x').unwrap() + current_index;
            let repeat_characters_count: usize =
                input[current_index + 1..separator_index].parse().unwrap();

            let closing_parentheses_index =
                input[current_index..].find(')').unwrap() + current_index;
            let repeat_times: u64 = input[separator_index + 1..closing_parentheses_index]
                .parse()
                .unwrap();

            let index_after_repeat = closing_parentheses_index + 1 + repeat_characters_count;

            let repeated_part = &input[closing_parentheses_index + 1 .. index_after_repeat];
            length += get_decompressed_length(repeated_part) * repeat_times;
            current_index = index_after_repeat;
        } else {
            current_index += 1;
            length += 1;
        }
    }

    length
}