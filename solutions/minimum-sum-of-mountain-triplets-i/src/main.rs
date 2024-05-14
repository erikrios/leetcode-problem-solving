fn main() {
    println!("{}", Solution::minimum_sum(vec![8,6,1,5,3]));
    println!("{}", Solution::minimum_sum(vec![5,4,8,7,10,2]));
    println!("{}", Solution::minimum_sum(vec![6,5,4,3,4,5]));
}

struct Solution;

impl Solution {
    pub fn minimum_sum(nums: Vec<i32>) -> i32 {
        let n = nums.len();

        let mut min_sum = i32::MAX;
        for i in 0..n-2 {
            for j in i+1..n-1 {
                for k in j+1..n {
                    let start = nums[i];
                    let middle = nums[j];
                    let end = nums[k];

                    if start < middle && middle > end {
                        let sum = start + middle + end;
                        if sum < min_sum {
                            min_sum = sum;
                        }
                    }
                }
            }
        }

        if min_sum == i32::MAX {
            return -1;
        }

        min_sum
    }
}
