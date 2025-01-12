fn main() {
    println!("{:?}", Solution::ways_to_split_array(vec![10, 4, -8, 7]));
    println!("{:?}", Solution::ways_to_split_array(vec![2, 3, 1, 0]));
}

struct Solution;

impl Solution {
    pub fn ways_to_split_array(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut prefix_sums = Vec::with_capacity(n);
        let mut sum = 0;

        for num in nums {
            sum += num as i64;
            prefix_sums.push(sum);
        }

        let mut res = 0;
        for i in 0..n - 1 {
            if prefix_sums[i] >= (prefix_sums[n - 1] - prefix_sums[i]) {
                res += 1;
            }
        }

        res
    }
}
