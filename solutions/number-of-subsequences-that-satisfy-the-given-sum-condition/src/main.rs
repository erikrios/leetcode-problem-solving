fn main() {
    println!("{}", Solution::num_subseq(vec![3, 5, 6, 7], 9));
    println!("{}", Solution::num_subseq(vec![3, 3, 6, 8], 10));
    println!("{}", Solution::num_subseq(vec![2, 3, 3, 4, 6, 7], 12));
    println!("{}", Solution::num_subseq(vec![1], 1));
    println!(
        "{}",
        Solution::num_subseq(
            vec![
                14, 4, 6, 6, 20, 8, 5, 6, 8, 12, 6, 10, 14, 9, 17, 16, 9, 7, 14, 11, 14, 15, 13,
                11, 10, 18, 13, 17, 17, 14, 17, 7, 9, 5, 10, 13, 8, 5, 18, 20, 7, 5, 5, 15, 19, 14
            ],
            22
        )
    );
}

struct Solution;

impl Solution {
    pub fn num_subseq(nums: Vec<i32>, target: i32) -> i32 {
        const MOD: i32 = 1_000_000_007;
        let mut nums = nums;
        nums.sort();
        let mut ans = 0;

        let mut l = 0;
        let mut r = nums.len() as i32 - 1;

        let mut power_of_2 = vec![1; nums.len()];
        for i in 1..nums.len() {
            power_of_2[i] = (power_of_2[i - 1] * 2) % MOD;
        }

        while l <= r {
            if nums[l as usize] + nums[r as usize] <= target {
                ans = (ans + power_of_2[(r - l) as usize]) % MOD;
                l += 1;
            } else {
                r -= 1;
            }
        }

        ans
    }
}
