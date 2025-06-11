use std::collections::HashMap;

fn main() {
    println!("{}", Solution::max_difference("aaaaabbc".to_string()));
    println!("{}", Solution::max_difference("abcabcab".to_string()));
    println!("{}", Solution::max_difference("tzt".to_string()));
}

struct Solution;

impl Solution {
    pub fn max_difference(s: String) -> i32 {
        let mut mapper = HashMap::with_capacity(26);

        for ch in s.into_bytes() {
            mapper
                .entry(ch)
                .and_modify(|counter| *counter += 1)
                .or_insert(1);
        }

        let mut odd_max = i32::MIN;
        let mut even_min = i32::MAX;

        for (_, counter) in mapper {
            if counter & 1 == 0 {
                if counter < even_min {
                    even_min = counter;
                }
            } else if counter > odd_max {
                odd_max = counter;
            }
        }

        odd_max - even_min
    }
}
