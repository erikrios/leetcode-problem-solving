fn main() {
    println!("{}", Solution::min_changes("1001".to_string()));
    println!("{}", Solution::min_changes("10".to_string()));
    println!("{}", Solution::min_changes("0000".to_string()));
}

struct Solution;

impl Solution {
    pub fn min_changes(s: String) -> i32 {
        let n = s.len();
        let s = s.into_bytes();
        let mut changes = 0;

        for i in (0..n).step_by(2) {
            let cur = s[i];
            let next = s[i + 1];

            if cur != next {
                changes += 1;
            }
        }

        changes
    }
}
