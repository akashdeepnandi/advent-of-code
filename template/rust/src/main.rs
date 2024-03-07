use std::env;
use std::fs::read_to_string;

fn main() {
    let lines: Vec<String> = read_to_string(env::args().nth(1).expect("Error: Missing file name"))
        .unwrap()
        .lines()
        .map(String::from)
        .collect();
    solve(lines);
}

fn solve(lines: Vec<String>) {
    for line in lines {
        println!("{}", line)
    }
}
