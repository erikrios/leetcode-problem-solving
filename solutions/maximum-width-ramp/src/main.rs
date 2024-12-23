fn main() {
    println!("{}", Solution::max_width_ramp(vec![6, 0, 8, 2, 1, 5]));
    println!(
        "{}",
        Solution::max_width_ramp(vec![9, 8, 1, 0, 1, 9, 4, 0, 4, 1])
    );
    println!("{}", Solution::max_width_ramp(vec![5, 4, 3, 2, 1]));
}

struct Solution;

impl Solution {
    pub fn max_width_ramp(nums: Vec<i32>) -> i32 {
        let mut stack = Vec::new();

        for (i, &num) in nums.iter().enumerate() {
            if let Some(&index) = stack.last() {
                if num <= nums[index] {
                    stack.push(i);
                }
            } else {
                stack.push(i);
            }
        }

        let mut max_ramp = 0;
        'outer: for (j, &num) in nums.iter().enumerate().rev() {
            while let Some(&index) = stack.last() {
                if num >= nums[index] {
                    max_ramp = max_ramp.max(j as i32 - index as i32);
                    stack.pop();
                } else {
                    continue 'outer;
                }
            }
            break;
        }

        max_ramp
    }
}
