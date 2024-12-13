fn main() {
    println!("{}", Solution::semi_ordered_permutation(vec![2, 1, 4, 3]));
    println!("{}", Solution::semi_ordered_permutation(vec![2, 4, 1, 3]));
    println!(
        "{}",
        Solution::semi_ordered_permutation(vec![1, 3, 4, 2, 5])
    );
    println!("{}", Solution::semi_ordered_permutation(vec![2, 1]));
    println!("{}", Solution::semi_ordered_permutation(vec![3, 2, 1]));
}

struct Solution;

impl Solution {
    pub fn semi_ordered_permutation(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut nums = nums;

        let mut min_num_operations = 0;
        while !(nums[0] == 1 && nums[n - 1] == n as i32) {
            for i in 0..n {
                let num = nums[i];
                if num == 1 && i > 0 {
                    let temp = nums[i - 1];
                    nums[i - 1] = num;
                    nums[i] = temp;
                    min_num_operations += 1;
                } else if num == n as i32 && i < n - 1 {
                    let temp = nums[i + 1];
                    nums[i + 1] = num;
                    nums[i] = temp;
                    min_num_operations += 1;
                }
            }
        }

        min_num_operations
    }
}
