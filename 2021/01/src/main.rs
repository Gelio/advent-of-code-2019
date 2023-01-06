use std::io::{stdin, Read};

fn main() {
    let mut buf = String::new();
    stdin().lock().read_to_string(&mut buf).unwrap();

    let input: Vec<_> = buf
        .trim()
        .lines()
        .map(|x| x.parse::<usize>().unwrap())
        .collect();

    let res: u32 = input
        .windows(2)
        .map(|elems| if elems[0] < elems[1] { 1 } else { 0 })
        .sum();
    println!("{res}");
}
