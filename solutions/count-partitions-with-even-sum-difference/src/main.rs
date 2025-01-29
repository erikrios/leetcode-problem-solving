fn main() {
    println!("{}", Solution::count_partitions(vec![10, 10, 3, 7, 6]));
    println!("{}", Solution::count_partitions(vec![1, 2, 2]));
    println!("{}", Solution::count_partitions(vec![2, 4, 6, 8]));
    println!("{}", Solution::count_partitions(vec![2, 2]));
}

struct Solution;

impl Solution {
    pub fn count_partitions(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut prefix_sum = Vec::with_capacity(n);
        let mut sum = 0;

        for num in nums {
            sum += num;
            prefix_sum.push(sum);
        }

        let mut counter = 0;
        for i in 1..n {
            let left_side = prefix_sum[i - 1];
            let right_side = prefix_sum[n - 1] - left_side;
            if (right_side - left_side).abs() & 1 == 0 {
                counter += 1;
            }
        }

        counter
    }
}
