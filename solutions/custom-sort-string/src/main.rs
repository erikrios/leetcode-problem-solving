use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::custom_sort_string("cba".to_string(), "abcd".to_string())
    );
    println!(
        "{}",
        Solution::custom_sort_string("bcafg".to_string(), "abcd".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn custom_sort_string(order: String, s: String) -> String {
        let mut s_trackers = HashMap::new();
        let n = s_trackers.len();

        for ch in s.into_bytes() {
            s_trackers.entry(ch).and_modify(|x| *x += 1).or_insert(1);
        }

        let mut results = String::with_capacity(n);

        for ch in order.into_bytes() {
            if let Some(&count) = s_trackers.get(&ch) {
                for _ in 0..count {
                    results.push(ch as char);
                }
                s_trackers.remove(&ch);
            }
        }

        for (ch, count) in s_trackers {
            for _ in 0..count {
                results.push(ch as char);
            }
        }

        results
    }
}
