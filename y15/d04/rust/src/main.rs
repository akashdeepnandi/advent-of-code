use std::env;
use std::fs::read_to_string;

fn main() {
    let lines: Vec<String> = read_to_string(env::args().nth(1).expect("Error: Missing file name"))
        .unwrap()
        .lines()
        .map(String::from)
        .collect();
    solve(lines.clone(), "00000");
    solve(lines, "000000");
}

fn solve(lines: Vec<String>, prefix: &str) {
    let input = lines.first().expect("Input Missing");
    let mut i = 0;
    let mut digest;
    loop {
        digest = md5::compute(input.to_owned() + &i.to_string());
        let digest_str = format!("{:x}", digest);
        if digest_str.starts_with(prefix) {
            break;
        }
        i += 1;
    }
    println!("{:x} Count - {}", digest, i + 1);
}
