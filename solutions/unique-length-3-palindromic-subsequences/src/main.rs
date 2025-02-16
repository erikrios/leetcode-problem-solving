use std::collections::{HashMap, HashSet};

fn main() {
    println!(
        "{}",
        Solution::count_palindromic_subsequence("aabca".to_string())
    );
    println!(
        "{}",
        Solution::count_palindromic_subsequence("adc".to_string())
    );
    println!(
        "{}",
        Solution::count_palindromic_subsequence("bbcbaba".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn count_palindromic_subsequence(s: String) -> i32 {
        let s = s.into_bytes();
        let mut hash_map = HashMap::new();

        for (i, &ch) in s.iter().enumerate() {
            hash_map
                .entry(ch)
                .and_modify(|indices: &mut Vec<usize>| {
                    indices.push(i);
                })
                .or_insert_with(|| {
                    let indices = vec![i];
                    indices
                });
        }

        let mut res = 0;
        for (_, indices) in hash_map {
            let n = indices.len();
            let start = indices[0];
            let end = indices[n - 1];

            if end - start <= 1 {
                continue;
            }

            let mut hash_set = HashSet::new();

            for &ch in s.iter().take(end).skip(start + 1) {
                hash_set.insert(ch);
            }

            res += hash_set.len();
        }

        res as i32
    }
}
