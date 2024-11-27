use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::num_of_pairs(
            vec![
                "777".to_string(),
                "7".to_string(),
                "77".to_string(),
                "77".to_string()
            ],
            "7777".to_string()
        )
    );
    println!(
        "{}",
        Solution::num_of_pairs(
            vec![
                "123".to_string(),
                "4".to_string(),
                "12".to_string(),
                "34".to_string()
            ],
            "1234".to_string()
        )
    );
    println!(
        "{}",
        Solution::num_of_pairs(
            vec!["1".to_string(), "1".to_string(), "1".to_string()],
            "11".to_string()
        )
    );
}

struct Solution;

impl Solution {
    pub fn num_of_pairs(nums: Vec<String>, target: String) -> i32 {
        let mut counter = 0;
        let mut hash_map = HashMap::new();

        for num in nums {
            if target.starts_with(&num) {
                let sub = &target[num.len()..];
                if let Some(count) = hash_map.get(sub) {
                    counter += count;
                }
            }

            if target.ends_with(&num) {
                let sub = &target[..target.len() - num.len()];
                if let Some(count) = hash_map.get(sub) {
                    counter += count;
                }
            }

            hash_map
                .entry(num)
                .and_modify(|count| *count += 1)
                .or_insert(1);
        }

        counter
    }
}
