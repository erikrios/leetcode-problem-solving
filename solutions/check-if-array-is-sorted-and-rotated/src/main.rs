fn main() {
    println!("{}", Solution::check(vec![3, 4, 5, 1, 2]));
    println!("{}", Solution::check(vec![2, 1, 3, 4]));
    println!("{}", Solution::check(vec![1, 2, 3]));
}

struct Solution;

impl Solution {
    pub fn check(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut counter = 0_u8;

        for i in 0..n {
            let num = nums[i];
            let front = if i == n - 1 { nums[0] } else { nums[i + 1] };

            if num > front {
                counter += 1;
                if counter == 2 {
                    return false;
                }
            }
        }

        true
    }
}
