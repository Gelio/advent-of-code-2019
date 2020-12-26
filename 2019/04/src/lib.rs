pub fn valid_passwords_count(min: u32, max: u32, predicate: fn(pass: u32) -> bool) -> u32 {
    (min..=max).filter(|&x| predicate(x)).count() as u32
}

struct DigitGroup {
    digit: u8,
    len: usize,
}

fn parse_digit_groups(pass: u32) -> Vec<DigitGroup> {
    let s = format!("{}", pass);
    let mut digit_groups: Vec<DigitGroup> = Vec::new();

    s.bytes().for_each(|digit| {
        match digit_groups.last_mut() {
            Some(dg) if dg.digit == digit => {
                dg.len += 1;
            }
            _ => {
                digit_groups.push(DigitGroup { digit, len: 1 });
            }
        };
    });

    digit_groups
}

pub fn is_valid_password_a(pass: u32) -> bool {
    let digit_groups = parse_digit_groups(pass);

    let not_increasing = digit_groups
        .iter()
        .fold((true, 0u8), |(r, p), c| (r && p <= c.digit, c.digit))
        .0;

    let has_double = digit_groups.iter().any(|dg| dg.len > 2);

    has_double && not_increasing
}

pub fn is_valid_password_b(pass: u32) -> bool {
    let digit_groups = parse_digit_groups(pass);

    let not_increasing = digit_groups
        .iter()
        .fold((true, 0u8), |(r, p), c| (r && p <= c.digit, c.digit))
        .0;

    let has_exactly_double_digits = digit_groups.iter().any(|dg| dg.len == 2);

    has_exactly_double_digits && not_increasing
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn valid_passwords_a() {
        assert!(is_valid_password_a(111111));
        assert!(!is_valid_password_a(223450));
        assert!(!is_valid_password_a(123789));
    }

    #[test]
    fn valid_passwords_b() {
        assert!(is_valid_password_b(112233));
        assert!(!is_valid_password_b(123444));
        assert!(is_valid_password_b(111122));
    }
}
