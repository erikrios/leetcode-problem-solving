use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::longest_square_streak(vec![4, 3, 6, 16, 8, 2])
    );
    println!("{}", Solution::longest_square_streak(vec![2, 3, 5, 6, 7]));
    println!(
        "{}",
        Solution::longest_square_streak(vec![4, 3, 4, 6, 16, 8, 2])
    );
    println!("{}", Solution::longest_square_streak(vec![3, 9, 81, 6561]));
}

struct Solution;

impl Solution {
    pub fn longest_square_streak(nums: Vec<i32>) -> i32 {
        let mut square_map = HashMap::new();
        let mut nums = nums;
        nums.sort();

        for num in nums {
            let square_val = (num as f64).sqrt();
            if square_val.fract() != 0.0 {
                square_map.insert(num, 1_usize);
            } else {
                let square_val = square_val as i32;
                if let Some(val) = square_map.get(&square_val) {
                    square_map.insert(num, val + 1);
                } else {
                    square_map.insert(num, 1);
                }
            }
        }

        let mut longest = 0;

        for &val in square_map.values() {
            longest = longest.max(val);
        }

        if longest < 2 {
            return -1;
        }

        longest as i32
    }
}
