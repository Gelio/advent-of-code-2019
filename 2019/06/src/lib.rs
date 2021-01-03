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

pub struct OrbitsGraph<'a> {
    // An edge exists from the center (the object being orbitted) to the object that orbits
    // e.g. COM)B
    // edges["COM"] = vec!["B"]
    edges: HashMap<&'a str, Vec<&'a str>>,
    reverse_edges: HashMap<&'a str, Vec<&'a str>>,
}

impl<'a> OrbitsGraph<'a> {
    pub fn parse(lines: impl Iterator<Item = &'a str>) -> Result<OrbitsGraph<'a>, ParseError<'a>> {
        let mut orbits_graph = OrbitsGraph {
            edges: HashMap::new(),
            reverse_edges: HashMap::new(),
        };

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
                .edges
                .entry(from)
                .and_modify(|v| v.push(to))
                .or_insert_with(|| vec![to]);

            orbits_graph
                .reverse_edges
                .entry(to)
                .and_modify(|v| v.push(from))
                .or_insert_with(|| vec![from]);
        });

        Ok(orbits_graph)
    }

    // Answer for part A. Counts the number of stars that orbit another star
    pub fn get_total_orbits(&self, star: &str, distance_from_center: i32) -> i32 {
        let default_orbiting_stars = Vec::new();
        let orbiting_stars = self.edges.get(star).unwrap_or(&default_orbiting_stars);

        let nested_orbits: i32 = orbiting_stars
            .into_iter()
            .map(|star| self.get_total_orbits(*star, distance_from_center + 1))
            .sum();

        distance_from_center + nested_orbits
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn error_when_invalid_orbits() {
        let input = vec!["a)b", "not a valid orbit"];

        let res = OrbitsGraph::parse(input.into_iter());

        assert!(res.is_err());
    }

    #[test]
    fn correctly_parses_orbits_graph() {
        let input = vec!["a)b", "b)c", "a)c"];

        let res = OrbitsGraph::parse(input.into_iter()).expect("cannot parse orbits");

        assert_eq!(res.edges.len(), 2);
        assert_eq!(res.edges["a"], vec!["b", "c"]);
        assert_eq!(res.edges["b"], vec!["c"]);
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

        let orbits_graph = OrbitsGraph::parse(input.into_iter()).expect("cannot parse orbits");
        let total_orbits = orbits_graph.get_total_orbits("COM", 0);

        assert_eq!(total_orbits, 42);
    }
}
