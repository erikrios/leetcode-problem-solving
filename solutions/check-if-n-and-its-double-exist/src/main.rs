use std::collections::HashSet;

fn main() {
    println!("{}", Solution::check_if_exist(vec![10, 2, 5, 3]));
    println!("{}", Solution::check_if_exist(vec![3, 1, 7, 11]));
    println!("{}", Solution::check_if_exist(vec![7, 1, 14, 11]));
}

struct Solution;

impl Solution {
    pub fn check_if_exist(arr: Vec<i32>) -> bool {
        let mut hash_set = HashSet::new();

        for num in arr {
            if hash_set.contains(&(num * 2)) || (num % 2 == 0 && hash_set.contains(&(num / 2))) {
                return true;
            }
            hash_set.insert(num);
        }

        false
    }
}
