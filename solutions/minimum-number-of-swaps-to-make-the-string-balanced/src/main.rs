fn main() {
    println!("{}", Solution::min_swaps("][][".to_string()));
    println!("{}", Solution::min_swaps("]]][[[".to_string()));
    println!("{}", Solution::min_swaps("[]".to_string()));
}

struct Solution;

impl Solution {
    pub fn min_swaps(s: String) -> i32 {
        let mut current_closing_bracket_counter = 0;
        let mut max_current_closing_bracket_counter = 0;

        for ch in s.into_bytes() {
            if ch == b'[' {
                current_closing_bracket_counter -= 1;
            } else {
                current_closing_bracket_counter += 1;
            }

            max_current_closing_bracket_counter =
                max_current_closing_bracket_counter.max(current_closing_bracket_counter);
        }

        (max_current_closing_bracket_counter + 1) / 2
    }
}
