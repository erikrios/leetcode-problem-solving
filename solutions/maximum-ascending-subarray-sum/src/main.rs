fn main() {
    println!(
        "{}",
        Solution::max_ascending_sum(vec![10, 20, 30, 5, 10, 50])
    );
    println!("{}", Solution::max_ascending_sum(vec![10, 20, 30, 40, 50]));
    println!(
        "{}",
        Solution::max_ascending_sum(vec![12, 17, 15, 13, 10, 11, 12])
    );
    println!("{}", Solution::max_ascending_sum(vec![1]));
}

struct Solution;

impl Solution {
    pub fn max_ascending_sum(nums: Vec<i32>) -> i32 {
        let mut prev = 0;
        let mut sum = 0;
        let mut max = 0;

        for num in nums {
            if num <= prev {
                sum = 0;
            }
            sum += num;
            max = max.max(sum);
            prev = num;
        }

        max
    }
}
