use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::maximum_subarray_sum(vec![1, 5, 4, 2, 9, 9, 9], 3)
    );
    println!("{}", Solution::maximum_subarray_sum(vec![4, 4, 4], 3));
}

struct Solution;

impl Solution {
    pub fn maximum_subarray_sum(nums: Vec<i32>, k: i32) -> i64 {
        let n = nums.len();

        let mut start_window: usize = 0;
        let mut window_sum: i64 = 0;
        let mut max_sum: i64 = 0;

        let mut mapper = HashMap::new();

        for end_window in 0..n {
            let num = nums[end_window];
            window_sum += num as i64;
            *mapper.entry(num).or_insert(0) += 1;
            if end_window >= (k - 1) as usize {
                if mapper.len() == k as usize && window_sum > max_sum {
                    max_sum = window_sum;
                }
                window_sum -= nums[start_window] as i64;
                *mapper.get_mut(&nums[start_window]).unwrap() -= 1;
                if *mapper.get_mut(&nums[start_window]).unwrap() == 0 {
                    mapper.remove(&(&nums[start_window]));
                }
                start_window += 1;
            }
        }

        max_sum
    }
}
