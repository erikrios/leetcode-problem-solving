fn main() {
    println!(
        "{}",
        Solution::remove_occurrences("daabcbaabcbc".to_string(), "abc".to_string())
    );
    println!(
        "{}",
        Solution::remove_occurrences("axxxxyyyyb".to_string(), "xy".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn remove_occurrences(s: String, part: String) -> String {
        let mut s = s;
        let mut prev = String::new();

        while prev != s {
            prev = s.clone();
            s = s.replacen(&part, "", 1);
        }

        s
    }
}
