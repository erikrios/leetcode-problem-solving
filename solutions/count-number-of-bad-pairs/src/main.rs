use std::collections::HashMap;

fn main() {
    println!("{}", Solution::count_bad_pairs(vec![4, 1, 3, 3]));
    println!("{}", Solution::count_bad_pairs(vec![1, 2, 3, 4, 5]));
}

struct Solution;

impl Solution {
    pub fn count_bad_pairs(nums: Vec<i32>) -> i64 {
        let mut total_pairs = 0;
        let mut good_pairs = 0;
        let mut counts = HashMap::new();

        for (i, num) in nums.into_iter().enumerate() {
            total_pairs += i as i64;
            let slope = num - i as i32;
            if let Some(&val) = counts.get(&slope) {
                good_pairs += val as i64;
            }
            counts.entry(slope).and_modify(|e| *e += 1).or_insert(1);
        }

        total_pairs - good_pairs
    }
}
