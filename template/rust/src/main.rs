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
    for line in lines {
        println!("{}", line)
    }
}
