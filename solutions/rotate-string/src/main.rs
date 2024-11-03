fn main() {
    println!(
        "{}",
        Solution::rotate_string("abcde".to_string(), "cdeab".to_string())
    );
    println!(
        "{}",
        Solution::rotate_string("abcde".to_string(), "abced".to_string())
    );
    println!(
        "{}",
        Solution::rotate_string("abcde".to_string(), "abcde".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn rotate_string(s: String, goal: String) -> bool {
        let mut s_temp = s.clone();

        while s_temp != goal {
            let ch = s_temp.remove(0);
            s_temp.push(ch);

            if s_temp == s {
                return false;
            }
        }

        true
    }
}
