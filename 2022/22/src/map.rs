use std::cmp::{max, min};

use crate::map_tiles::{MapTiles, Tile};

#[derive(Debug)]
pub struct MapLine {
    start_index: usize,
    end_index: usize,
    wall_indices: Vec<usize>,
}

impl MapLine {
    fn row(tiles: &Vec<Tile>) -> Self {
        let start_index = tiles
            .iter()
            .position(|tile| *tile != Tile::Empty)
            .expect("Row is empty");

        let end_index = tiles
            .iter()
            .rposition(|tile| *tile != Tile::Empty)
            .expect("Row is empty");

        // TODO: use start_index and end_index for the range
        let wall_indices: Vec<_> = (0..tiles.len())
            .filter(|index| tiles[*index] == Tile::Wall)
            .collect();

        Self {
            start_index,
            end_index,
            wall_indices,
        }
    }

    fn column(rows: &Vec<Vec<Tile>>, index: usize) -> Self {
        let start_index = rows
            .iter()
            .position(|row| row[index] != Tile::Empty)
            .expect("Column is empty");

        let end_index = rows
            .iter()
            .rposition(|row| row[index] != Tile::Empty)
            .expect("Column is empty");

        let wall_indices: Vec<_> = (0..rows.len())
            .filter(|column| rows[*column][index] == Tile::Wall)
            .collect();

        Self {
            start_index,
            end_index,
            wall_indices,
        }
    }

    fn len(&self) -> usize {
        self.end_index - self.start_index + 1
    }

    pub fn move_positive(&self, from_index: usize, steps: usize) -> usize {
        assert!((self.start_index..=self.end_index).contains(&from_index));

        if self.wall_indices.is_empty() {
            return (from_index + steps) % self.len();
        }

        // NOTE: there is at least 1 wall in the line
        // Let's try to find it without wrapping around

        let wall_index_without_wrapping = match self.wall_indices.binary_search(&from_index) {
            Ok(_) => panic!("Index {from_index} is a wall"),
            Err(wall_internal_index) => {
                if wall_internal_index == self.wall_indices.len() {
                    // NOTE: there is no wall in the positive direction.
                    None
                } else {
                    Some(self.wall_indices[wall_internal_index])
                }
            }
        };
        if let Some(wall_index_without_wrapping) = wall_index_without_wrapping {
            return min(from_index + steps, wall_index_without_wrapping - 1);
        }

        // NOTE: no wall in the positive direction.
        // Check if wrapping is needed
        if from_index + steps <= self.end_index {
            return from_index + steps;
        }

        // NOTE: wrap around until the first wall
        let first_wall_index = self.wall_indices[0];

        if first_wall_index == self.start_index {
            // NOTE: there is a wall at the front of the line.
            // The max we can go is the end of the line.
            self.end_index
        } else {
            // The "- 1" is for the step from the end tile to the start tile
            let steps_left_after_wrapping = steps - (self.end_index - from_index) - 1;

            min(
                self.start_index + steps_left_after_wrapping,
                first_wall_index - 1,
            )
        }
    }

    pub fn move_negative(&self, from_index: usize, steps: usize) -> usize {
        assert!((self.start_index..=self.end_index).contains(&from_index));

        if self.wall_indices.is_empty() {
            return (self.len() + from_index - (steps % self.len())) % self.len();
        }

        // NOTE: there is at least 1 wall in the line
        // Let's try to find it without wrapping around

        let wall_index_without_wrapping = match self.wall_indices.binary_search(&from_index) {
            Ok(_) => panic!("Index {from_index} is a wall"),
            Err(wall_internal_index) => {
                if wall_internal_index == 0 {
                    // NOTE: there is no wall in the negative direction.
                    None
                } else {
                    Some(self.wall_indices[wall_internal_index - 1])
                }
            }
        };
        if let Some(wall_index_without_wrapping) = wall_index_without_wrapping {
            return max::<isize>(
                from_index as isize - steps as isize,
                wall_index_without_wrapping as isize + 1,
            ) as usize;
        }

        // NOTE: no wall in the negative direction.
        // Check if wrapping is needed
        if steps <= from_index - self.start_index {
            return from_index - steps;
        }

        // NOTE: wrap around until the first wall
        let last_wall_index = *self.wall_indices.last().expect(
            "There are no walls. We checked earlier that there should be at least one wall.",
        );

        if last_wall_index == self.end_index {
            // NOTE: there is a wall at the end of the line.
            // The max we can go is the start of the line.
            self.start_index
        } else {
            // The "- 1" is for the step from the start tile to the end tile
            let steps_left_after_wrapping =
                (steps - (from_index - self.start_index) - 1) % self.len();

            max(
                self.end_index - steps_left_after_wrapping,
                last_wall_index + 1,
            )
        }
    }

    pub fn first_open_tile(&self) -> Option<usize> {
        (self.start_index..=self.end_index).find(|index| !self.wall_indices.contains(index))
    }
}

#[derive(Debug)]
pub struct Map {
    rows: Vec<MapLine>,
    columns: Vec<MapLine>,
}

impl Map {
    pub fn new(tiles: MapTiles) -> Self {
        Map {
            rows: tiles.iter().map(MapLine::row).collect(),
            columns: (0..tiles[0].len())
                .map(|index| MapLine::column(&tiles, index))
                .collect(),
        }
    }

    pub fn rows(&self) -> &Vec<MapLine> {
        &self.rows
    }

    pub fn columns(&self) -> &Vec<MapLine> {
        &self.columns
    }
}

#[cfg(test)]
mod tests {
    use nom::combinator::all_consuming;

    use crate::map_tiles::parser::map_tiles;

    use super::*;

    fn get_map_line(input: &str) -> MapLine {
        let map = Map::new(
            all_consuming(map_tiles)(input)
                .expect("Could not parse map tiles")
                .1,
        );

        map.rows.into_iter().next().expect("Map rows are empty")
    }

    #[test]
    fn computes_first_empty_tile() {
        assert_eq!(get_map_line("......#....").first_open_tile(), Some(0));
        assert_eq!(get_map_line("###...#....").first_open_tile(), Some(3));
        assert_eq!(get_map_line("##").first_open_tile(), None);
    }

    #[test]
    fn moves_forward() {
        let map_line = get_map_line("......#....");

        assert_eq!(
            map_line.move_positive(map_line.start_index, 1),
            map_line.start_index + 1
        );
        assert_eq!(
            map_line.move_positive(map_line.start_index, 5),
            map_line.start_index + 5
        );
        assert_eq!(
            map_line.move_positive(map_line.start_index, 6),
            // NOTE: hits a wall
            map_line.start_index + 5
        );
        assert_eq!(
            map_line.move_positive(map_line.start_index, 50),
            // NOTE: hits a wall
            map_line.start_index + 5
        );
    }

    #[test]
    fn moves_forward_no_walls() {
        let map_line = get_map_line("....");

        assert_eq!(
            map_line.move_positive(map_line.start_index, 1),
            map_line.start_index + 1
        );
        assert_eq!(
            map_line.move_positive(map_line.start_index, 3),
            map_line.start_index + 3
        );
        assert_eq!(
            map_line.move_positive(map_line.start_index, 6),
            map_line.start_index + 2
        );
    }

    #[test]
    fn moves_forward_wrapping_around() {
        let row = get_map_line("..#....#....");

        assert_eq!(row.move_positive(row.end_index - 1, 1), row.end_index);
        assert_eq!(row.move_positive(row.end_index, 1), 0);
        assert_eq!(row.move_positive(row.end_index, 2), 1);
        assert_eq!(
            row.move_positive(row.end_index, 5),
            // NOTE: hits a wall
            1
        );
    }

    #[test]
    fn moves_forward_wrapping_around_with_wall_at_start() {
        let row = get_map_line("#.#....#....");

        assert_eq!(row.move_positive(row.end_index, 1), row.end_index);
        assert_eq!(row.move_positive(row.end_index - 1, 2), row.end_index);
        assert_eq!(row.move_positive(row.end_index - 2, 1), row.end_index - 1);
    }
    #[test]
    fn moves_back() {
        let map_line = get_map_line("..#...");

        assert_eq!(
            map_line.move_negative(map_line.end_index, 1),
            map_line.end_index - 1
        );
        assert_eq!(
            map_line.move_negative(map_line.end_index, 2),
            map_line.end_index - 2
        );
        assert_eq!(
            map_line.move_negative(map_line.end_index, 6),
            // NOTE: hits a wall
            map_line.end_index - 2
        );
        assert_eq!(
            map_line.move_negative(map_line.end_index, 50),
            // NOTE: hits a wall
            map_line.end_index - 2
        );
    }

    #[test]
    fn moves_back_no_walls() {
        let input = "....";
        let map_line = get_map_line(input);

        assert_eq!(
            map_line.move_negative(map_line.end_index, 1),
            map_line.end_index - 1
        );
        assert_eq!(
            map_line.move_negative(map_line.end_index, 3),
            map_line.end_index - 3
        );
        assert_eq!(
            map_line.move_negative(map_line.end_index, 6),
            map_line.start_index + 1
        );
        assert_eq!(
            map_line.move_negative(map_line.start_index, 41),
            map_line.end_index
        );
    }

    #[test]
    fn moves_back_wrapping_around() {
        let row = get_map_line("..#....#..");

        assert_eq!(row.move_negative(row.start_index + 1, 2), row.end_index);
        assert_eq!(row.move_negative(row.start_index + 1, 1), row.start_index);
        assert_eq!(row.move_negative(row.start_index, 2), row.end_index - 1);
        assert_eq!(
            row.move_negative(row.start_index, 5),
            // NOTE: hits a wall
            row.end_index - 1
        );
    }

    #[test]
    fn moves_back_wrapping_around_with_wall_at_end() {
        let row = get_map_line("...#....#...#");

        assert_eq!(row.move_negative(row.start_index, 1), row.start_index);
        assert_eq!(row.move_negative(row.start_index + 1, 2), row.start_index);
        assert_eq!(row.move_negative(row.start_index + 2, 50), row.start_index);
    }
}
