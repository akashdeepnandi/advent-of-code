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
    let mut total = 0;
    let mut actual_total = 0;
    let mut extra_total = 0;

    for line in lines {
        let len = line.len();
        let mut actual = 0;

        let chars: Vec<char> = line.chars().collect();
        let mut i = 1;
        while i < chars.len() - 1 {
            if chars[i] == '\\' {
                match chars[i + 1] {
                    '\\' | '"' => {
                        actual += 1;
                        i += 2;
                    }
                    'x' => {
                        actual += 1;
                        i += 4;
                    }
                    _ => {
                        actual += 1;
                        i += 1;
                    }
                }
            } else {
                actual += 1;
                i += 1;
            }
        }

        let mut i = 0;
        let mut extra = 2;
        while i < chars.len() {
            match chars[i] {
                '\\' | '"' => {
                    extra += 2;
                }
                _ => {
                    extra += 1;
                }
            }
            i += 1;
        }

        total += len;
        actual_total += actual;
        extra_total += extra
    }

    let part1 = total - actual_total;
    println!("Part 1: {}", part1);

    let part2 = extra_total - total;
    println!("Part 2: {}", part2)
}
