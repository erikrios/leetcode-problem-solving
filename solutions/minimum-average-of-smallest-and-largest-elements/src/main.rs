fn main() {
    println!(
        "{}",
        Solution::minimum_average(vec![7, 8, 3, 4, 15, 13, 4, 1])
    );
    println!("{}", Solution::minimum_average(vec![1, 9, 8, 3, 10, 5]));
    println!("{}", Solution::minimum_average(vec![1, 2, 3, 7, 8, 9]));
}

struct Solution;

impl Solution {
    pub fn minimum_average(nums: Vec<i32>) -> f64 {
        let mut nums = nums;
        nums.sort();
        let mut min_avg = f64::MAX;

        for i in 0..nums.len() / 2 {
            let start = nums[i];
            let end = nums[nums.len() - 1 - i];

            let avg = (start + end) as f64 / 2.0;
            min_avg = min_avg.min(avg);
        }

        min_avg
    }
}
