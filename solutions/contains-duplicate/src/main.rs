use std::collections::HashSet;

fn main() {
    println!("{}", Solution::contains_duplicate(vec![1, 2, 3, 1]));
    println!("{}", Solution::contains_duplicate(vec![1, 2, 3, 4]));
    println!(
        "{}",
        Solution::contains_duplicate(vec![1, 1, 1, 3, 3, 4, 3, 2, 4, 2])
    );
}

struct Solution;

impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut hash_set = HashSet::new();

        for num in nums {
            if hash_set.contains(&num) {
                return true;
            }

            hash_set.insert(num);
        }

        false
    }
}
