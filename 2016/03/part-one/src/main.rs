use std::io;

fn main() {
    let mut line = String::new();
    let mut triangles_count = 0;

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

        let mut numbers = line.split_whitespace().map(|stringified_number| i32::from_str_radix(stringified_number, 10).unwrap());
        let a = numbers.next().unwrap();
        let b = numbers.next().unwrap();
        let c = numbers.next().unwrap();

        if is_triangle(a, b, c) {
            triangles_count += 1;
        }
    }

    println!("Triangles: {}", triangles_count);
}

fn is_triangle(a: i32, b: i32, c: i32) -> bool {
    a + b > c && a + c > b && b + c > a
}