fn main() {
    println!("{}", Solution::reduction_operations(vec![5, 1, 3]));
    println!("{}", Solution::reduction_operations(vec![1, 1, 1]));
    println!("{}", Solution::reduction_operations(vec![1, 1, 2, 2, 3]));
}

struct Solution;

impl Solution {
    pub fn reduction_operations(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut nums = nums;
        nums.sort();

        let mut max = nums[n - 1];
        let mut operation_num = 0;

        for i in (0..n).rev() {
            let num = nums[i];
            if num < max {
                max = num;
                operation_num += n - i - 1;
            }
        }

        operation_num as i32
    }
}
