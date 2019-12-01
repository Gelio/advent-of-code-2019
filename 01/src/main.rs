use std::io;

fn main() {
    let mut sum: i32 = 0;

    loop {
        let mut line = String::new();
        let result = io::stdin().read_line(&mut line);
        if result.is_err() {
            break;
        }

        let trimmed_line = line.trim();
        if trimmed_line.is_empty() {
            break;
        }

        let mass = trimmed_line.parse::<i32>().expect("Not a valid number");
        sum += get_fuel_required(mass);
    }

    println!("Sum: {}", sum);
}

fn get_fuel_required(mass: i32) -> i32 {
    return mass / 3 - 2;
}
