fn main() {
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("52233".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("5494".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("1111111".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("4411794".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("00".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("0".to_string())
    );
    println!(
        "{}",
        Solution::longest_semi_repetitive_substring("24489929009".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn longest_semi_repetitive_substring(s: String) -> i32 {
        let s: Vec<char> = s.chars().collect();
        let mut window_start = 0;
        let mut longest = 0;
        let mut dup_counter = 0;

        for window_end in 0..s.len() {
            if window_end > 0 {
                let end = s[window_end];
                let prev_end = s[window_end - 1];
                if prev_end == end {
                    dup_counter += 1;
                }
            }
            while dup_counter == 2 {
                window_start += 1;
                let start = s[window_start];
                if window_start > 0 {
                    let prev_start = s[window_start - 1];
                    if prev_start == start {
                        dup_counter -= 1;
                    }
                }
            }
            longest = std::cmp::max(longest, (window_end - window_start) as i32 + 1);
        }

        longest
    }
}
