use std::collections::HashMap;

fn main() {
    println!("{}", Solution::minimum_length("abaacbcbb".to_string()));
    println!("{}", Solution::minimum_length("aa".to_string()));
    println!("{}", Solution::minimum_length("aaaaaaaaa".to_string()));
}

struct Solution;

impl Solution {
    pub fn minimum_length(s: String) -> i32 {
        let s = s.into_bytes();
        let mut mapper = HashMap::new();
        let mut result = s.len();

        for ch in s {
            mapper.entry(ch).and_modify(|x| *x += 1).or_insert(1);
        }

        for (_, occ) in mapper {
            if occ & 1 == 1 {
                result -= occ - 1;
            } else {
                result -= occ - 2;
            }
        }

        result as i32
    }
}
