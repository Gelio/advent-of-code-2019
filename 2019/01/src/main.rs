use std::io;

fn main() {
    let masses = parse_numbers().expect("Cannot parse numbers");

    let sum: i64 = masses.into_iter().map(get_total_fuel_required).sum();

    println!("Sum: {}", sum);
}

fn parse_numbers() -> Result<Vec<i64>, String> {
    let mut numbers: Vec<i64> = Vec::new();

    loop {
        let mut line = String::new();
        match io::stdin().read_line(&mut line) {
            Ok(0) => break,
            Ok(_) => match line.trim().parse() {
                Ok(number) => numbers.push(number),
                Err(e) => return Err(format!("cannot parse line {}: {}", line, e.to_string())),
            },
            Err(e) => return Err(format!("error when reading line: {}", e.to_string())),
        }
    }

    Ok(numbers)
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
