fn main() {
    println!("{}", Solution::min_max_game(vec![1, 3, 5, 2, 4, 8, 2, 2]));
    println!("{}", Solution::min_max_game(vec![3]));
    println!("{}", Solution::min_max_game(vec![93, 40]));
}

struct Solution;

impl Solution {
    pub fn min_max_game(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        if n == 1 {
            return nums[0];
        }

        let mut new_nums = Vec::with_capacity(n / 2);

        for i in 0..n / 2 {
            if i & 1 == 0 {
                let min = nums[2 * i].min(nums[2 * i + 1]);
                new_nums.push(min);
            } else {
                let max = nums[2 * i].max(nums[2 * i + 1]);
                new_nums.push(max);
            }
        }

        Self::min_max_game(new_nums)
    }
}
