fn main() {
    println!("{}", Solution::count_max_or_subsets(vec![3, 1]));
    println!("{}", Solution::count_max_or_subsets(vec![2, 2, 2]));
    println!("{}", Solution::count_max_or_subsets(vec![3, 2, 1, 5]));
}

struct Solution;

impl Solution {
    pub fn count_max_or_subsets(nums: Vec<i32>) -> i32 {
        let max = nums.iter().fold(0, |acc, &num| acc | num);
        Self::backtracking(0, 0, 0, max, &nums) as i32
    }

    fn backtracking(i: usize, total: usize, bitwise_res: i32, target: i32, nums: &[i32]) -> usize {
        let n = nums.len();

        if i == n {
            if total > 0 && bitwise_res == target {
                return 1;
            } else {
                return 0;
            }
        }

        if bitwise_res > target {
            return 0;
        }

        // Not take
        let mut count = Self::backtracking(i + 1, total, bitwise_res, target, nums);

        let bitwise_res = bitwise_res | nums[i];

        // Take
        count += Self::backtracking(i + 1, total + 1, bitwise_res, target, nums);

        count
    }
}
