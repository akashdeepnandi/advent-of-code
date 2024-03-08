use regex::Regex;
use std::fs::read_to_string;
use std::{char, env};

fn main() {
    let lines: Vec<String> = read_to_string(env::args().nth(1).expect("Error: Missing file name"))
        .unwrap()
        .lines()
        .map(String::from)
        .collect();
    solve(lines);
}

fn solve(lines: Vec<String>) {
    // PART 1
    let re_vowels = Regex::new(r"(a|e|i|o|u)").unwrap();
    let re_bad = Regex::new(r"(ab|cd|pq|xy)").unwrap();
    let mut count1 = 0;
    for line in &lines {
        let vowels = re_vowels.find_iter(line);
        if vowels.count() < 3 {
            continue;
        }

        // check for double
        let mut chars = line.chars();
        let mut check = chars.next();
        let mut doubles = false;
        for curc in chars {
            if check == Some(curc) {
                doubles = true;
                break;
            }
            check = Some(curc)
        }
        if !doubles {
            continue;
        }

        let bad = re_bad.find_iter(line);
        if bad.count() > 0 {
            continue;
        }

        count1 += 1;
    }

    println!("Part 1: {}", count1);

    // PART 2
    let mut count2 = 0;
    for line in &lines {
        let mut pairs = false;
        let chars: Vec<char> = line.chars().collect();

        for i in 0..chars.len() - 3 {
            for j in i + 2..chars.len() - 1 {
                if chars[i] == chars[j] && chars[i + 1] == chars[j + 1] {
                    pairs = true;
                    break;
                }
            }
        }

        if !pairs {
            continue;
        }

        let mut one_between = false;
        for i in 0..chars.len() - 2 {
            if chars[i] == chars[i + 2] {
                one_between = true;
                break;
            }
        }

        if !one_between {
            continue;
        }

        count2 += 1;
    }

    println!("Part 2 {}", count2);
}
