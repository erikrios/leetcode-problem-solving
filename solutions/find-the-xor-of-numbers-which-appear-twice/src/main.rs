use std::collections::HashSet;

fn main() {
    println!("{}", Solution::duplicate_numbers_xor(vec![1, 2, 1, 3]));
    println!("{}", Solution::duplicate_numbers_xor(vec![1, 2, 3]));
    println!("{}", Solution::duplicate_numbers_xor(vec![1, 2, 2, 1]));
}

struct Solution;

impl Solution {
    pub fn duplicate_numbers_xor(nums: Vec<i32>) -> i32 {
        let mut set = HashSet::new();
        let mut result = 0;

        for &num in nums.iter() {
            if let Some(_) = set.get(&num) {
                result ^= num
            } else {
                set.insert(num);
            }
        }

        result
    }
}
