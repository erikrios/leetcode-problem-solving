fn main() {
    println!(
        "{}",
        Solution::append_characters("coaching".to_string(), "coding".to_string())
    );
    println!(
        "{}",
        Solution::append_characters("abcde".to_string(), "a".to_string())
    );
    println!(
        "{}",
        Solution::append_characters("z".to_string(), "abcde".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn append_characters(s: String, t: String) -> i32 {
        let mut s_ptr = 0;
        let mut t_ptr = 0;

        let m = s.len();
        let n = t.len();

        let s: Vec<char> = s.chars().collect();
        let t: Vec<char> = t.chars().collect();

        while s_ptr < m && t_ptr < n {
            if s[s_ptr] == t[t_ptr] {
                t_ptr += 1;
            }
            s_ptr += 1;
        }

        (n - t_ptr) as i32
    }
}
