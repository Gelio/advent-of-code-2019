use std::{error::Error, fmt::Display};

#[derive(Debug)]
pub struct ParseError {
    invalid_char: String,
    wrapped_error: Box<dyn Error>,
}

impl Error for ParseError {}

impl Display for ParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "cannot parse {}: {}",
            self.invalid_char, self.wrapped_error
        )
    }
}

pub fn parse_from_string(s: &str) -> Result<Vec<isize>, Vec<ParseError>> {
    let (numbers, errors): (Vec<_>, Vec<_>) = s
        .trim()
        .split(',')
        .map(|c| {
            c.parse::<isize>().map_err(|e| ParseError {
                invalid_char: c.to_owned(),
                wrapped_error: Box::new(e),
            })
        })
        .partition(Result::is_ok);

    if !errors.is_empty() {
        return Err(errors.into_iter().map(Result::unwrap_err).collect());
    }

    Ok(numbers.into_iter().map(Result::unwrap).collect())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn valid_program() {
        let input = "3,8,1001,8,10,8";
        let program = parse_from_string(input);

        assert!(program.is_ok(), "parsing failed");
        assert_eq!(
            program.unwrap(),
            vec![3, 8, 1001, 8, 10, 8],
            "invalid contents"
        );
    }

    #[test]
    fn invalid_program() {
        let input = "3,8,abc,8,10,8";
        let program = parse_from_string(input);

        assert!(
            program.is_err(),
            "parsing successful despite invalid number"
        );

        let errors = program.unwrap_err();
        assert_eq!(errors.len(), 1, "invalid number of errors");
        assert_eq!(
            errors[0].invalid_char, "abc",
            "incorrectly matched invalid char"
        );
    }
}
