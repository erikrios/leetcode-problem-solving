fn main() {
    println!(
        "{}",
        Solution::max_distinct_elements(vec![1, 2, 2, 3, 3, 4], 2)
    );
    println!("{}", Solution::max_distinct_elements(vec![4, 4, 4, 4], 1));
}

struct Solution;

impl Solution {
    pub fn max_distinct_elements(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        nums.sort();
        let n = nums.len();
        let mut diff_counter = 0;

        for i in 0..n {
            let num = nums[i];
            if i == 0 {
                nums[i] -= k;
                diff_counter += 1;
                continue;
            }
            let prev = nums[i - 1];
            let mut min_addition = prev - num + 1;

            if min_addition < -k {
                min_addition = -k;
            } else if min_addition > k {
                min_addition = k;
            }

            nums[i] += min_addition;

            if nums[i] != prev {
                diff_counter += 1;
            }
        }

        diff_counter
    }
}
