use std::{collections::HashMap, error::Error, fmt::Display, vec};

#[derive(Debug)]
pub enum OrbitsParseError<'a> {
    InvalidLine(&'a str),
    MultipleOrbits {
        body: &'a str,
        first_orbit: &'a str,
        second_orbit: &'a str,
    },
}

impl<'a> Error for OrbitsParseError<'a> {}

impl<'a> Display for OrbitsParseError<'a> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::InvalidLine(line) => {
                write!(f, "cannot parse line {}", line)
            }
            Self::MultipleOrbits {
                body,
                first_orbit,
                second_orbit,
            } => {
                write!(
                    f,
                    "two or more orbits found for {}: {}, {}",
                    body, first_orbit, second_orbit
                )
            }
        }
    }
}

#[derive(Debug)]
pub struct OrbitPathError {
    center: String,
    destination: String,
}

impl Error for OrbitPathError {}

impl Display for OrbitPathError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "cannot find path from {} to destination {}",
            self.center, self.destination
        )
    }
}

pub struct OrbitsGraph<'a> {
    // An edge exists from the center (the object being orbitted) to the object that orbits
    // e.g. COM)B
    // edges["COM"] = vec!["B"]
    edges: HashMap<&'a str, Vec<&'a str>>,
    // Edge to a body that is closer to the center
    reverse_edge: HashMap<&'a str, &'a str>,
}

impl<'a> OrbitsGraph<'a> {
    pub fn parse(
        lines: impl Iterator<Item = &'a str>,
    ) -> Result<OrbitsGraph<'a>, OrbitsParseError<'a>> {
        let mut orbits_graph = OrbitsGraph {
            edges: HashMap::new(),
            reverse_edge: HashMap::new(),
        };

        let parsed_lines: Vec<_> = lines
            .map(|line| {
                let parts: Vec<_> = line.split(')').collect();
                if parts.len() != 2 {
                    Err(OrbitsParseError::InvalidLine(line))
                } else {
                    Ok((parts[0], parts[1]))
                }
            })
            .collect::<Result<_, _>>()?;

        parsed_lines
            .into_iter()
            .map(|(from, to)| {
                orbits_graph
                    .edges
                    .entry(from)
                    .and_modify(|v| v.push(to))
                    .or_insert_with(|| vec![to]);

                if let Some(existing_orbit) = orbits_graph.reverse_edge.insert(to, from) {
                    return Err(OrbitsParseError::MultipleOrbits {
                        body: from,
                        first_orbit: to,
                        second_orbit: existing_orbit,
                    });
                }

                Ok(())
            })
            .collect::<Result<_, _>>()?;

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

    pub fn get_path(
        &self,
        center: &'a str,
        destination: &'a str,
    ) -> Result<Vec<&str>, OrbitPathError> {
        let mut path = Vec::new();
        let mut current = destination;

        while current != center {
            path.push(current);
            current = self.reverse_edge[current];
        }
        path.push(center);

        path.reverse();

        Ok(path)
    }

    pub fn get_distance(
        &self,
        center: &str,
        node1: &str,
        node2: &str,
    ) -> Result<usize, OrbitPathError> {
        let path1 = self.get_path(center, node1)?;
        let path2 = self.get_path(center, node2)?;

        let common_path_len = get_common_prefix_len(&path1, &path2);

        Ok(path1.len() + path2.len() - 2 * common_path_len - 2)
    }
}

fn get_common_prefix_len<T: PartialEq>(v1: &[T], v2: &[T]) -> usize {
    v1.iter()
        .zip(v2.iter())
        .enumerate()
        .find(|(_, (a, b))| a != b)
        .map_or(v1.len() - 1, |(len, _)| len)
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
        let input = vec!["a)b", "b)c", "a)d"];

        let res = OrbitsGraph::parse(input.into_iter()).expect("cannot parse orbits");

        assert_eq!(res.edges.len(), 2);
        assert_eq!(res.edges["a"], vec!["b", "d"]);
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

    #[test]
    fn gets_path_between_center_and_some_body() {
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
        K)L
        K)YOU
        I)SAN"
            .lines()
            .map(|line| line.trim());

        let graph = OrbitsGraph::parse(input.into_iter()).expect("cannot parse orbits");

        let path = graph.get_path("COM", "SAN").unwrap();

        assert_eq!(path, vec!["COM", "B", "C", "D", "I", "SAN"]);
    }

    #[test]
    fn gets_common_prefix_length_when_exists_and_same_length() {
        let v1 = vec!["a", "b", "c"];
        let v2 = vec!["a", "b", "d"];

        let res = get_common_prefix_len(&v1, &v2);

        assert_eq!(res, 2);
    }

    #[test]
    fn gets_common_prefix_length_when_exists_and_different_length() {
        let v1 = vec!["a", "b", "c"];
        let v2 = vec!["a", "b"];

        let res = get_common_prefix_len(&v1, &v2);

        assert_eq!(res, 2);
    }

    #[test]
    fn reports_common_prefix_not_found() {
        let v1 = vec!["a", "b", "c"];
        let v2 = vec!["b", "c"];

        let res = get_common_prefix_len(&v1, &v2);

        assert_eq!(res, 0);
    }

    #[test]
    fn correctly_gets_distance_between_2_stars() {
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
        K)L
        K)YOU
        I)SAN"
            .lines()
            .map(|line| line.trim());

        let graph = OrbitsGraph::parse(input.into_iter()).expect("cannot parse orbits");

        let distance = graph.get_distance("COM", "SAN", "YOU").unwrap();

        assert_eq!(distance, 4);
    }
}
