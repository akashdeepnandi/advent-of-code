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
    let mut wrapping_paper = 0;
    let mut rybbon = 0;

    for line in lines {
        let dims: Vec<i32> = line
            .split('x')
            .filter_map(|d| d.parse::<i32>().ok())
            .collect();

        let (l, w, h) = (dims[0], dims[1], dims[2]);
        // part 1
        wrapping_paper += 2 * (l * w + w * h + h * l) + (l * w).min(w * h).min(h * l);

        // part 2
        rybbon += (2 * (l + w)).min(2 * (w + h)).min(2 * (h + l)) + l * w * h;
    }

    println!("Part 1: {}", wrapping_paper);
    println!("Part 2: {}", rybbon);
}
