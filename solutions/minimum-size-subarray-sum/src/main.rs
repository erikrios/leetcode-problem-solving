fn main() {
    println!("{}", Solution::min_sub_array_len(7, vec![2, 3, 1, 2, 4, 3]));
    println!("{}", Solution::min_sub_array_len(4, vec![1, 4, 4]));
    println!(
        "{}",
        Solution::min_sub_array_len(11, vec![1, 1, 1, 1, 1, 1, 1, 1])
    );
}

struct Solution;

impl Solution {
    pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut start_window = 0;
        let mut min_len = i32::MAX;
        let mut sum_window = 0;

        for end_window in 0..n {
            sum_window += nums[end_window];

            while sum_window >= target {
                let len = (end_window - start_window + 1) as i32;
                if len < min_len {
                    min_len = len;
                }
                sum_window -= nums[start_window];
                start_window += 1;
            }
        }

        if min_len == i32::MAX {
            return 0;
        }

        min_len
    }
}
