use aoc_2019_04;

fn main() {
    let min = 128392;
    let max = 643281;

    println!(
        "Result A: {}",
        aoc_2019_04::valid_passwords_count(min, max, aoc_2019_04::is_valid_password_a)
    );
    println!(
        "Result B: {}",
        aoc_2019_04::valid_passwords_count(min, max, aoc_2019_04::is_valid_password_b)
    );
}
