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

fn look_and_say(sequence: &[char]) -> Vec<char> {
    let mut result: Vec<char> = Vec::new();
    let mut current_digit: char = sequence[0];
    let mut count: usize = 0;

    for &d in sequence {
        if current_digit == d {
            count += 1;
        } else {
            result.extend(count.to_string().chars());
            result.push(current_digit);

            current_digit = d;
            count = 1
        }
    }
    //add final
    result.extend(count.to_string().chars());
    result.push(current_digit);
    return result;
}
fn evolve_sequence(mut sequence: Vec<char>, generation: usize) -> Vec<char> {
    for _ in 0..generation {
        sequence = look_and_say(&sequence);
    }

    sequence
}

fn solve(lines: Vec<String>) {
    let input: Vec<char> = lines.iter().next().unwrap().chars().collect();

    let part1: Vec<char> = evolve_sequence(input.clone(), 40);
    println!("Part 1: {}", part1.len());
    let part2: Vec<char> = evolve_sequence(input, 50);
    println!("Part 2: {}", part2.len());
}
