mod helper;

use helper::read_lines;
use std::collections::{HashMap, HashSet};
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
    let moves: Vec<char> = lines.first().unwrap().chars().collect();
    let mut dir: HashMap<char, (i32, i32)> = HashMap::new();
    dir.insert('^', (-1, 0));
    dir.insert('v', (1, 0));
    dir.insert('<', (0, -1));
    dir.insert('>', (0, 1));

    // FOR PART 1
    let mut positions: HashSet<(i32, i32)> = HashSet::new();
    let mut santa_pos = (0, 0);
    for m in &moves {
        positions.insert(santa_pos);
        santa_pos = (
            santa_pos.0 + dir.get(m).unwrap().0,
            santa_pos.1 + dir.get(m).unwrap().1,
        );
    }
    println!("Part 1: {}", positions.len());

    // FOR PART 2
    let mut positions: HashSet<(i32, i32)> = HashSet::new();
    let mut both_pos = [(0, 0), (0, 0)];

    for (i, m) in moves.iter().enumerate() {
        positions.insert(both_pos[i % 2]);
        both_pos[i % 2] = (
            both_pos[i % 2].0 + dir.get(m).unwrap().0,
            both_pos[i % 2].1 + dir.get(m).unwrap().1,
        );
    }

    println!("Part 2: {}", positions.len());
}
