extern crate md5;

use std::io;

fn main() {
    let mut line = String::new();

    if let Err(error) = io::stdin().read_line(&mut line) {
        println!("Error: {}", error);
        std::process::exit(1);
    }

    let line = line.trim();
    if line.len() == 0 {
        println!("Empty line");
        std::process::exit(1);
    }

    let mut hash_source = String::from(line);

    let mut code: Vec<char> = Vec::new();
    let mut number = 0;
    let mut debug_index = 0;

    while code.len() < 8 {
        debug_index += 1;
        if debug_index == 200000 {
            debug_index = 0;
            println!("Number: {}", number);    
        }

        hash_source.drain(line.len()..);
        hash_source.push_str(&number.to_string());

        let hash = md5::compute(hash_source.as_bytes());
        let mut hash_valid = true;

        for index in 0..2 {
            if hash[index] != 0 {
                hash_valid = false;
                break;
            }
        }
        if hash[2] & 0xf0 != 0 {
            hash_valid = false;
        }

        number = number + 1;
        if !hash_valid {
            continue;
        }
        code.push(hex_byte_to_string(hash[2] & 0x0f));
        println!("Found: {} -> {}", hash[2] & 0x0f, code[code.len() - 1]);
    }

    println!("Code: {:?}", code);
}

fn hex_byte_to_string(byte: u8) -> char {
    if byte < 10 {
        return (('0' as u8) + byte) as char
    }

    (('a' as u8) + (byte - 10)) as char
}
