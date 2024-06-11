fn main() {
    println!(
        "{}",
        Solution::find_max_average(vec![1, 12, -5, -6, 50, 3], 4)
    );
    println!("{}", Solution::find_max_average(vec![5], 1));
}

struct Solution;

impl Solution {
    pub fn find_max_average(nums: Vec<i32>, k: i32) -> f64 {
        let mut start_window: usize = 0;
        let mut window_sum = 0;
        let mut max_average = f64::MIN;

        let n = nums.len();

        for end_window in 0..n {
            let num = nums[end_window];
            window_sum += num;

            if end_window >= (k - 1) as usize {
                let average = window_sum as f64 / k as f64;
                if average > max_average {
                    max_average = average;
                }
                window_sum -= nums[start_window];
                start_window += 1;
            }
        }

        max_average
    }
}
