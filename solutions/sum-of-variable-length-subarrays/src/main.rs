fn main() {
    println!("{}", Solution::subarray_sum(vec![2, 3, 1]));
    println!("{}", Solution::subarray_sum(vec![3, 1, 1, 2]));
}

struct Solution;

impl Solution {
    pub fn subarray_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut prefix_sum = Vec::with_capacity(n);

        let mut sum = 0;
        let mut sub_sum = 0;
        for (i, num) in nums.into_iter().enumerate() {
            sum += num;
            prefix_sum.push(sum);

            let start = 0.max(i as i32 - num) as usize;

            let sub = sum - if start > 0 { prefix_sum[start - 1] } else { 0 };
            sub_sum += sub;
        }

        sub_sum
    }
}
