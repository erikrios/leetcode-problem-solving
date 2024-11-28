fn main() {
    println!("{}", Solution::minimum_steps("101".to_string()));
    println!("{}", Solution::minimum_steps("100".to_string()));
    println!("{}", Solution::minimum_steps("0111".to_string()));
    println!("{}", Solution::minimum_steps("100100100".to_string()));
}

struct Solution;

impl Solution {
    pub fn minimum_steps(s: String) -> i64 {
        let mut j = -1;
        let mut swap_count = 0;
        for (i, ch) in s.bytes().enumerate() {
            if ch == b'0' && j >= 0 {
                let range = i - j as usize;
                swap_count += range as i64;
                j += 1;
            } else if ch == b'1' && j == -1 {
                j = i as i32;
            }
        }

        swap_count
    }
}
