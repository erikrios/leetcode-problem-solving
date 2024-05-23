fn main() {
    println!("{}", Solution::is_array_special(vec![1]));
    println!("{}", Solution::is_array_special(vec![2, 1, 4]));
    println!("{}", Solution::is_array_special(vec![4, 3, 1, 6]));
}

struct Solution;

impl Solution {
    pub fn is_array_special(nums: Vec<i32>) -> bool {
        for i in 0..nums.len() - 1 {
            let cur = nums[i];
            let next = nums[i + 1];

            if cur % 2 == 0 && next % 2 == 0 {
                return false;
            }

            if cur % 2 != 0 && next % 2 != 0 {
                return false;
            }
        }

        true
    }
}
