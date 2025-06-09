fn main() {
    println!("{}", Solution::clear_stars("aaba*".to_string()));
    println!("{}", Solution::clear_stars("abc".to_string()));
}

struct Solution {}

impl Solution {
    pub fn clear_stars(s: String) -> String {
        let mut stack = Vec::new();

        for ch in s.chars() {
            if ch != '*' {
                stack.push(ch);
            } else {
                if let Some(pos) = (0..stack.len()).rev().min_by_key(|&i| stack[i]) {
                    stack.remove(pos);
                }
            }
        }

        stack.iter().collect()
    }
}
