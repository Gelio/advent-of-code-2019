use std::{collections::HashMap, error::Error, fmt::Display};

#[derive(Debug)]
pub struct ParseError<'a> {
    invalid_line: &'a str,
}

impl<'a> Error for ParseError<'a> {}

impl<'a> Display for ParseError<'a> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "parse error, cannot parse line {}", self.invalid_line)
    }
}

type OrbitsGraph<'a> = HashMap<&'a str, Vec<&'a str>>;

pub fn parse_orbits<'a>(
    lines: impl Iterator<Item = &'a str>,
) -> Result<OrbitsGraph<'a>, ParseError<'a>> {
    let mut orbits_graph: OrbitsGraph = HashMap::new();

    let parsed_lines: Vec<_> = lines
        .map(|line| {
            let parts: Vec<_> = line.split(')').collect();
            if parts.len() != 2 {
                Err(ParseError { invalid_line: line })
            } else {
                Ok((parts[0], parts[1]))
            }
        })
        .collect::<Result<_, _>>()?;

    parsed_lines.into_iter().for_each(|(from, to)| {
        orbits_graph
            .entry(from)
            .and_modify(|v| v.push(to))
            .or_insert_with(|| vec![to]);
    });

    Ok(orbits_graph)
}

// Answer for part A. Counts the number of stars that orbit another star
pub fn get_total_orbits<'a>(
    orbits_graph: &OrbitsGraph<'a>,
    star: &str,
    distance_from_center: i32,
) -> i32 {
    let default_orbiting_stars = Vec::new();
    let orbiting_stars = orbits_graph.get(star).unwrap_or(&default_orbiting_stars);

    let nested_orbits: i32 = orbiting_stars
        .into_iter()
        .map(|star| get_total_orbits(orbits_graph, *star, distance_from_center + 1))
        .sum();

    distance_from_center + nested_orbits
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn error_when_invalid_orbits() {
        let input = vec!["a)b", "not a valid orbit"];

        let res = parse_orbits(input.into_iter());

        assert!(res.is_err());
    }

    #[test]
    fn correctly_parses_orbits_graph() {
        let input = vec!["a)b", "b)c", "a)c"];

        let res = parse_orbits(input.into_iter()).expect("cannot parse orbits");

        assert_eq!(res.len(), 2);
        assert_eq!(res["a"], vec!["b", "c"]);
        assert_eq!(res["b"], vec!["c"]);
    }

    #[test]
    fn computes_total_orbits() {
        let input = "COM)B
        B)C
        C)D
        D)E
        E)F
        B)G
        G)H
        D)I
        E)J
        J)K
        K)L"
        .lines()
        .map(|line| line.trim());

        let orbit_graph = parse_orbits(input.into_iter()).expect("cannot parse orbits");
        let total_orbits = get_total_orbits(&orbit_graph, "COM", 0);

        assert_eq!(total_orbits, 42);
    }
}
