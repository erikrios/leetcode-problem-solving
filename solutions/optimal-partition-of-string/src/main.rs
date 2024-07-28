use std::collections::HashSet;

fn main() {
    println!("{}", Solution::partition_string("abacaba".to_string()));
    println!("{}", Solution::partition_string("ssssss".to_string()));
}

struct Solution;

impl Solution {
    pub fn partition_string(s: String) -> i32 {
        let mut counter = 0;
        let mut hash_set = HashSet::new();

        for ch in s.chars() {
            if hash_set.contains(&ch) {
                counter += 1;
                hash_set.clear();
            }

            hash_set.insert(ch);
        }

        counter + 1
    }
}
