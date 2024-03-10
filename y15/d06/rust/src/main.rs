use std::fs::read_to_string;
use std::{cmp, env};

fn main() {
    let lines: Vec<String> = read_to_string(env::args().nth(1).expect("Error: Missing file name"))
        .unwrap()
        .lines()
        .map(String::from)
        .collect();
    solve(lines);
}

fn solve(lines: Vec<String>) {
    enum Instruction {
        TurnOn,
        TurnOff,
        Toggle,
    }
    // Part 1
    let mut lights = vec![vec![0; 1000]; 1000];
    for line in &lines {
        let parts: Vec<&str> = line.split(' ').collect();
        let (command, start, end) = match parts[1] {
            "off" => (Instruction::TurnOff, parts[2], parts[4]),
            "on" => (Instruction::TurnOn, parts[2], parts[4]),
            _ => (Instruction::Toggle, parts[1], parts[3]),
        };

        let start: Vec<usize> = start
            .split(',')
            .map(|d| d.parse::<usize>().unwrap())
            .collect();
        let end: Vec<usize> = end
            .split(',')
            .map(|d| d.parse::<usize>().unwrap())
            .collect();

        for row in lights.iter_mut().take(end[0] + 1).skip(start[0]) {
            for elem in row.iter_mut().take(end[1] + 1).skip(start[1]) {
                *elem = match command {
                    Instruction::TurnOff => 0,
                    Instruction::TurnOn => 1,
                    Instruction::Toggle => match elem {
                        0 => 1,
                        1 => 0,
                        _ => 1,
                    },
                }
            }
        }
    }

    let part1: i32 = lights.iter().flat_map(|r| r.iter()).sum();

    println!("Part 1: {}", part1);
    // ----
    // Part 2
    let mut lights = vec![vec![0; 1000]; 1000];

    for line in &lines {
        let parts: Vec<&str> = line.split(' ').collect();
        let (command, start, end) = match parts[1] {
            "off" => (Instruction::TurnOff, parts[2], parts[4]),
            "on" => (Instruction::TurnOn, parts[2], parts[4]),
            _ => (Instruction::Toggle, parts[1], parts[3]),
        };

        let start: Vec<usize> = start
            .split(',')
            .map(|d| d.parse::<usize>().unwrap())
            .collect();
        let end: Vec<usize> = end
            .split(',')
            .map(|d| d.parse::<usize>().unwrap())
            .collect();

        for row in lights.iter_mut().take(end[0] + 1).skip(start[0]) {
            for elem in row.iter_mut().take(end[1] + 1).skip(start[1]) {
                *elem = match command {
                    Instruction::TurnOff => cmp::max(*elem - 1, 0),
                    Instruction::TurnOn => *elem + 1,
                    Instruction::Toggle => *elem + 2,
                }
            }
        }
    }

    let part2: i32 = lights.iter().flat_map(|r| r.iter()).sum();

    println!("Part 2: {}", part2);
}
