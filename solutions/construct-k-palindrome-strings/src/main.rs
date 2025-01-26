use std::collections::HashMap;

fn main() {
    println!("{}", Solution::can_construct("annabelle".to_string(), 2));
    println!("{}", Solution::can_construct("leetcode".to_string(), 3));
    println!("{}", Solution::can_construct("true".to_string(), 4));
}

struct Solution;

impl Solution {
    pub fn can_construct(s: String, k: i32) -> bool {
        if s.len() < k as usize {
            return false;
        }

        let mut s_mapper = HashMap::new();

        for ch in s.into_bytes() {
            s_mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1);
        }

        let mut odd_counter = 0;

        for (_, v) in s_mapper {
            if v & 1 == 1 {
                odd_counter += 1;
                if odd_counter > k {
                    return false;
                }
            }
        }

        true
    }
}
