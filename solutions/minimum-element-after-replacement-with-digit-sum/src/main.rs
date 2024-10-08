fn main() {
    println!("{}", Solution::min_element(vec![10, 12, 13, 14]));
    println!("{}", Solution::min_element(vec![1, 2, 3, 4]));
    println!("{}", Solution::min_element(vec![999, 19, 199]));
}

struct Solution;

impl Solution {
    pub fn min_element(nums: Vec<i32>) -> i32 {
        let mut min = 10_000;

        for mut num in nums {
            let mut sum = 0;
            while num > 0 {
                let digit = num % 10;
                num /= 10;
                sum += digit;
            }
            min = min.min(sum);
        }

        min
    }
}
