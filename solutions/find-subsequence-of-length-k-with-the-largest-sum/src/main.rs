fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    pub fn max_subsequence(nums: Vec<i32>, k: i32) -> Vec<i32> {
        vec![]
    }

    fn dfs(
        nums: &[i32],
        k: usize,
        i: usize,
        current_subsequence: &[i32],
        current_max_sum: &mut i32,
        current_max_subsequence: &[i32],
    ) {
        if current_subsequence.len() == k {
            let current_sum: i32 = current_subsequence.iter().sum();
            if current_sum > *current_max_sum {
                *current_max_sum = current_sum;
            }
        }

        for i in i + 1..=nums.len() - k {}
    }
}
