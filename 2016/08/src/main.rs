use std::io;

extern crate regex;
use regex::Regex;

fn main() {
    let mut line = String::new();

    let width = 50;
    let height = 6;
    let mut screen = prepare_screen(width, height);

    let rectangle_regex = Regex::new(r"rect (\d+)x(\d+)").unwrap();
    let rotate_row_regex = Regex::new(r"rotate row y=(\d+) by (\d+)").unwrap();
    let rotate_column_regex = Regex::new(r"rotate column x=(\d+) by (\d+)").unwrap();


    loop {
        line.clear();

        io::stdin()
            .read_line(&mut line)
            .expect("Error while reading the line");

        let line = line.trim();
        if line.len() == 0 {
            break;
        }

        for cap in rectangle_regex.captures_iter(line) {
            let rectangle_width = usize::from_str_radix(&cap[1], 10).unwrap();
            let rectangle_height = usize::from_str_radix(&cap[2], 10).unwrap();

            for x in 0..rectangle_width {
                for y in 0..rectangle_height {
                    screen[x][y] = true;
                }
            }
        }

        for cap in rotate_row_regex.captures_iter(line) {
            let row_number = usize::from_str_radix(&cap[1], 10).unwrap();
            let rotate_amount = i32::from_str_radix(&cap[2], 10).unwrap();

            let mut new_row = vec![false; width as usize];

            for x in 0..width {
                let new_index = (x + rotate_amount) % width;
                new_row[new_index as usize] = screen[x as usize][row_number];
            }

            for (index, _) in new_row.iter().enumerate() {
                screen[index][row_number] = new_row[index];
            }
        }

        for cap in rotate_column_regex.captures_iter(line) {
            let column_number = usize::from_str_radix(&cap[1], 10).unwrap();
            let rotate_amount = i32::from_str_radix(&cap[2], 10).unwrap();
            
            let mut new_column = vec![false; height as usize];

            for y in 0..height {
                let new_index = (y + rotate_amount) % height;
                new_column[new_index as usize] = screen[column_number][y as usize];
            }

            for (index, _) in new_column.iter().enumerate() {
                screen[column_number][index] = new_column[index];
            }
        }
    }

    let mut visible_signs = 0;

    print_screen(&screen);

    for column in screen.iter() {
        for element in column.iter() {
            if *element {
                visible_signs += 1;
            }
        }
    }

    println!("Visible: {}", visible_signs);

}

fn prepare_screen(width: i32, height: i32) -> Vec<Vec<bool>> {
    let mut screen = Vec::new();

    for _ in 0..width {
        let mut column = Vec::new();

        for _ in 0..height {
            column.push(false);
        }

        screen.push(column);
    }

    screen
}

fn print_screen(screen: &Vec<Vec<bool>>) {
    let width = screen.len();

    if width == 0 {
        return;
    }

    let height = screen[0].

    for y in 0..height {
        for x in 0..width {
            if screen[x][y] {
                print!("#");
            } else {
                print!(".");
            }
        }

        println!("");
    }
}