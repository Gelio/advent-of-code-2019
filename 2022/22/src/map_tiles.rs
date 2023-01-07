use std::ops::Deref;

#[derive(Debug, PartialEq, Clone)]
pub enum Tile {
    /// Tile does not exist (' ')
    Empty,
    /// Character can appear here ('.')
    Open,
    /// A wall ('#')
    Wall,
}

impl TryFrom<char> for Tile {
    type Error = String;

    fn try_from(value: char) -> Result<Self, Self::Error> {
        match value {
            ' ' => Ok(Tile::Empty),
            '.' => Ok(Tile::Open),
            '#' => Ok(Tile::Wall),
            c => Err(format!("{c} is not a valid tile")),
        }
    }
}

#[derive(Debug)]
pub struct MapTiles(Vec<Vec<Tile>>);

impl TryFrom<Vec<Vec<Tile>>> for MapTiles {
    type Error = String;

    fn try_from(value: Vec<Vec<Tile>>) -> Result<Self, Self::Error> {
        if value.is_empty() {
            return Err("Empty map".to_string());
        }

        let expected_columns = value[0].len();

        if let Some((index, row)) = value
            .iter()
            .enumerate()
            .find(|(_, row)| row.len() != expected_columns)
        {
            return Err(format!("Row with index {index} in the map has a different number of columns ({0}, expected {expected_columns})", row.len()));
        }

        Ok(MapTiles(value))
    }
}

impl Deref for MapTiles {
    type Target = Vec<Vec<Tile>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

pub mod parser {
    use nom::{
        character::complete::{anychar, newline},
        combinator::map_res,
        multi::{many1, separated_list1},
        IResult,
    };

    use super::*;

    fn tile(input: &str) -> IResult<&str, Tile> {
        map_res(anychar, TryInto::try_into)(input)
    }

    fn row(input: &str) -> IResult<&str, Vec<Tile>> {
        many1(tile)(input)
    }

    fn rows(input: &str) -> IResult<&str, Vec<Vec<Tile>>> {
        separated_list1(newline, row)(input)
    }

    pub fn map_tiles(input: &str) -> IResult<&str, MapTiles> {
        map_res(rows, |mut rows| {
            let columns = rows
                .iter()
                .map(|row| row.len())
                .max()
                .expect("Map has 0 rows");

            for row in rows.iter_mut() {
                row.resize(columns, Tile::Empty);
            }

            rows.try_into()
        })(input)
    }

    #[cfg(test)]
    mod tests {
        use nom::combinator::all_consuming;

        use super::*;

        #[test]
        fn parses_example_map() {
            let input = "        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.";

            let result = all_consuming(map_tiles)(input);
            result.expect("Parsing failed");
        }
    }
}
