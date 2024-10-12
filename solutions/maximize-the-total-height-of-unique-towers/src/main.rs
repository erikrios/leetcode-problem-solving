use std::collections::{BTreeMap, HashSet};

fn main() {
    println!("{}", Solution::maximum_total_sum(vec![2, 3, 4, 3]));
    println!("{}", Solution::maximum_total_sum(vec![15, 10]));
    println!("{}", Solution::maximum_total_sum(vec![2, 2, 1]));
    println!("{}", Solution::maximum_total_sum(vec![4, 6, 5, 1, 6]));
}

struct Solution;

impl Solution {
    pub fn maximum_total_sum(maximum_height: Vec<i32>) -> i64 {
        let mut b_tree_map = BTreeMap::new();

        for max_height in maximum_height {
            b_tree_map
                .entry(max_height)
                .and_modify(|x| *x += 1)
                .or_insert(1);
        }

        let mut hash_set = HashSet::new();
        let mut sum = 0_i64;
        let mut last_inserted = 0;

        for (mut max_height, count) in b_tree_map.into_iter().rev() {
            for _ in 0..count {
                if hash_set.contains(&max_height) {
                    max_height = last_inserted - 1;
                    if max_height == 0 {
                        return -1;
                    }
                }

                hash_set.insert(max_height);
                last_inserted = max_height;
                sum += max_height as i64;
            }
        }

        sum
    }
}
