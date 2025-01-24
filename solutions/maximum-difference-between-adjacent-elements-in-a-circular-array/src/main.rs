fn main() {
    println!("{}", Solution::max_adjacent_distance(vec![1, 2, 4]));
    println!("{}", Solution::max_adjacent_distance(vec![-5, -10, -5]));
}

struct Solution;

impl Solution {
    pub fn max_adjacent_distance(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut max_abs = 0;

        for i in 0..n {
            let cur = nums[i];
            let next = if i == n - 1 { nums[0] } else { nums[i + 1] };

            max_abs = max_abs.max((cur - next).abs());
        }

        max_abs
    }
}
