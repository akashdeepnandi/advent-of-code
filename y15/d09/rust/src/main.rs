use std::collections::HashMap;
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

fn permute<T: Clone>(arr: &mut [T], i: usize, result: &mut Vec<Vec<T>>) {
    if i == arr.len() {
        result.push(arr.to_vec());
        return;
    }

    for j in i..arr.len() {
        arr.swap(i, j);
        // Permute the subtree
        permute(arr, i + 1, result);
        // backtrack
        arr.swap(i, j);
    }
}

fn get_permutations<T: Clone>(arr: &mut [T]) -> Vec<Vec<T>> {
    let mut result: Vec<Vec<T>> = Vec::new();
    permute(arr, 0, &mut result);
    result
}

fn solve(lines: Vec<String>) {
    let mut graph: HashMap<String, HashMap<String, i32>> = HashMap::new();

    for line in lines {
        let parts: Vec<&str> = line.split_whitespace().collect();
        let u = parts[0];
        let v = parts[2];
        let d: i32 = parts[parts.len() - 1].parse().unwrap();

        graph
            .entry(u.to_string())
            .or_default()
            .insert(v.to_string(), d);
        graph
            .entry(v.to_string())
            .or_default()
            .insert(u.to_string(), d);
    }

    let mut nodes: Vec<String> = graph.keys().cloned().collect();

    let routes = get_permutations(&mut nodes);

    let distances: Vec<i32> = routes
        .iter()
        .map(|route| route.windows(2).map(|w| graph[&w[0]][&w[1]]).sum())
        .collect();

    let part1 = distances.iter().min().copied().unwrap();
    println!("Part 1: {}", part1);

    let part2 = distances.iter().max().copied().unwrap();
    println!("Part 2: {}", part2);
}
