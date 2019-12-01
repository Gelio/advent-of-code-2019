use std::io;

fn main() {
    let masses = parse_numbers();

    let sum: i64 = masses.into_iter().map(get_total_fuel_required).sum();

    println!("Sum: {}", sum);
}

fn parse_numbers() -> Vec<i64> {
    let mut numbers: Vec<i64> = Vec::new();

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

        let number = trimmed_line.parse::<i64>().expect("Not a valid number");
        numbers.push(number);
    }

    numbers
}

fn get_fuel_required(mass: i64) -> i64 {
    mass / 3 - 2
}

fn get_total_fuel_required(initial_mass: i64) -> i64 {
    let mut mass = initial_mass;
    let mut sum = 0;

    while mass > 0 {
        let fuel_for_mass = get_fuel_required(mass);
        if fuel_for_mass <= 0 {
            break;
        }
        sum += fuel_for_mass;
        mass = fuel_for_mass;
    }

    sum
}
