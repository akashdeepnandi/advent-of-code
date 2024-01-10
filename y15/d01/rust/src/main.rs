mod helper;

use helper::read_lines;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        let file = &args[1];
        println!("File: {}", file);
        let lines = read_lines(file);
        solve(lines);
    } else {
        println!("No file name provided");
    }
}

fn solve(lines: Vec<String>) {
    let input: Vec<char> = lines[0].chars().collect();

    let mut floor = 0;
    let mut basement_at = -1;
    for (i, c) in input.iter().enumerate() {
        if *c == '(' {
            floor += 1
        } else {
            floor -= 1
        }

        if basement_at == -1 && floor == -1 {
            basement_at = i as i32 + 1
        }
    }
    println!("{}", floor);
    println!("{}", basement_at);
}
