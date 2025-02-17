use std::collections::BinaryHeap;

fn main() {
    println!("{}", Solution::min_operations(vec![2, 11, 10, 1, 3], 10));
    println!("{}", Solution::min_operations(vec![1, 1, 2, 4, 9], 20));
    println!(
        "{}",
        Solution::min_operations(
            vec![1000000000, 999999999, 1000000000, 999999999, 1000000000, 999999999],
            1000000000
        )
    );
}

struct Solution;

impl Solution {
    pub fn min_operations(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums: BinaryHeap<i64> = nums.into_iter().map(|x| -(x as i64)).collect();
        let mut operation = 0;

        while nums.len() >= 2 && -*nums.peek().unwrap() < k as i64 {
            let first_smallest = -nums.pop().unwrap();
            let second_smallest = -nums.pop().unwrap();

            let new_num = first_smallest * 2 + second_smallest;

            nums.push(-new_num);

            operation += 1;
        }

        operation
    }
}
