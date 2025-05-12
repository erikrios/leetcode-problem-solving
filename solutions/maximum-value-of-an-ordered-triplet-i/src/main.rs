fn main() {
    println!("{}", Solution::maximum_triplet_value(vec![12, 6, 1, 2, 7]));
    println!("{}", Solution::maximum_triplet_value(vec![1, 10, 3, 4, 19]));
    println!("{}", Solution::maximum_triplet_value(vec![1, 2, 3]));
    println!(
        "{}",
        Solution::maximum_triplet_value(vec![1000000, 1, 1000000])
    );
}

struct Solution;

impl Solution {
    pub fn maximum_triplet_value(nums: Vec<i32>) -> i64 {
        let n = nums.len();
        let mut max = 0;

        for i in 0..n {
            for j in i + 1..n {
                for k in j + 1..n {
                    let res = (nums[i] as i64 - nums[j] as i64) * nums[k] as i64;
                    if res > max {
                        max = res;
                    }
                }
            }
        }

        max
    }
}
