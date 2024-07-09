fn main() {
    println!("{}", Solution::minimum_operations(vec![1, 2, 3, 4]));
    println!("{}", Solution::minimum_operations(vec![3, 6, 9]));
}

struct Solution;

impl Solution {
    pub fn minimum_operations(nums: Vec<i32>) -> i32 {
        let mut counter = 0;

        for num in nums {
            if num % 3 != 0 {
                counter += 1;
            }
        }

        counter
    }
}
