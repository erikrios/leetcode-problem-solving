fn main() {
    println!("{}", Solution::triangular_sum(vec![1, 2, 3, 4, 5]));
    println!("{}", Solution::triangular_sum(vec![5]));
}

struct Solution;

impl Solution {
    pub fn triangular_sum(nums: Vec<i32>) -> i32 {
        if nums.len() == 1 {
            return nums[0];
        }

        let mut nums = nums;

        for i in 0..nums.len() - 1 {
            nums[i] = (nums[i] + nums[i + 1]) % 10;
        }

        nums.pop();

        Self::triangular_sum(nums)
    }
}
