use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::max_equal_rows_after_flips(vec![vec![0, 1], vec![1, 1]])
    );
    println!(
        "{}",
        Solution::max_equal_rows_after_flips(vec![vec![0, 1], vec![1, 0]])
    );
    println!(
        "{}",
        Solution::max_equal_rows_after_flips(vec![vec![0, 0, 0], vec![0, 0, 1], vec![1, 1, 0]])
    );
}

struct Solution;

impl Solution {
    pub fn max_equal_rows_after_flips(matrix: Vec<Vec<i32>>) -> i32 {
        let mut hash_map = HashMap::new();
        let mut maximum = 0;

        for nums in matrix {
            let mut pattern = String::from("*");

            let mut prev = nums[0];
            for num in nums.into_iter().skip(1) {
                if num != prev {
                    pattern.push('|');
                }
                pattern.push('*');
                prev = num;
            }

            hash_map
                .entry(pattern.clone())
                .and_modify(|e| *e += 1)
                .or_insert(1);

            maximum = maximum.max(*hash_map.get(&pattern).unwrap())
        }

        maximum
    }
}
