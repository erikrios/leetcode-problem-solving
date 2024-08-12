fn main() {
    println!(
        "{}",
        Solution::final_position_of_snake(2, vec!["RIGHT".to_string(), "DOWN".to_string()])
    );
    println!(
        "{}",
        Solution::final_position_of_snake(
            3,
            vec!["DOWN".to_string(), "RIGHT".to_string(), "UP".to_string()]
        )
    );
}

struct Solution;

impl Solution {
    pub fn final_position_of_snake(n: i32, commands: Vec<String>) -> i32 {
        let mut i = 0;
        let mut j = 0;

        for command in commands {
            match command.as_str() {
                "UP" => i -= 1,
                "RIGHT" => j += 1,
                "DOWN" => i += 1,
                "LEFT" => j -= 1,
                _ => {}
            }
        }

        i * n + j
    }
}
