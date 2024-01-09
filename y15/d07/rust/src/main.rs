mod helper;

use helper::read_lines;
use std::{collections::HashMap, env};

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

fn get_wire_val(wire: &str, values: &HashMap<&str, i32>) -> i32 {
    match wire.parse::<i32>() {
        Ok(v) => v,
        Err(_) => match values.get(wire) {
            Some(v) => *v,
            None => -1,
        },
    }
}

fn solve(lines: Vec<String>) {
    let mut instructions: Vec<[String; 4]> = Vec::new();
    for line in lines {
        if let Some((opr, tid)) = line.split_once(" -> ") {
            let operation: Vec<&str> = opr.split(' ').collect();
            let ol = operation.len();
            let gate = match ol {
                2 => "NOT",
                3 => operation[1],
                _ => "",
            }
            .to_string();
            let op1 = match ol {
                2 => operation[1],
                _ => operation[0],
            }
            .to_string();
            let op2 = match ol {
                3 => operation[2],
                _ => "",
            }
            .to_string();

            instructions.push([tid.to_string(), gate, op1, op2])
        }
    }

    let part1 = get_value(&instructions, "a");
    println!("Part 1: {}", part1);

    for ins in &mut instructions {
        if ins[0] == "b" {
            ins[2] = part1.to_string();
            break;
        }
    }
    let part2 = get_value(&instructions, "a");
    println!("Part 2: {}", part2);
}

fn get_value(instructions: &Vec<[String; 4]>, target: &str) -> i32 {
    let mut values: HashMap<&str, i32> = HashMap::new();
    loop {
        let mut resolved = true;
        for [tid, gate, op1, op2] in instructions {
            let v1 = get_wire_val(op1, &values);
            if v1 == -1 {
                resolved = false;
                continue;
            }

            if gate.is_empty() {
                values.insert(tid, v1);
            } else if gate == "NOT" {
                values.insert(tid, 65535 - v1);
            } else {
                let v2 = get_wire_val(op2, &values);
                if v2 == -1 {
                    resolved = false;
                    continue;
                }
                match gate.as_str() {
                    "AND" => values.insert(tid, v1 & v2),
                    "OR" => values.insert(tid, v1 | v2),
                    "LSHIFT" => values.insert(tid, v1 << v2 & 0xffff),
                    _ => values.insert(tid, v1 >> v2),
                };
            };
        }

        if resolved {
            break;
        }
    }
    *values.get(target).unwrap_or(&-1)
}
