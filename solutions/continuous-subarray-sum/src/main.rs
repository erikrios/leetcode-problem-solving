use std::collections::HashMap;

fn main() {
    println!("{}", Solution::check_subarray_sum(vec![23, 2, 4, 6, 7], 6));
    println!("{}", Solution::check_subarray_sum(vec![23, 2, 6, 4, 7], 6));
    println!("{}", Solution::check_subarray_sum(vec![23, 2, 6, 4, 7], 13));
    println!("{}", Solution::check_subarray_sum(vec![0, 0], 1));
}

struct Solution;

impl Solution {
    pub fn check_subarray_sum(nums: Vec<i32>, k: i32) -> bool {
        let n = nums.len();

        if n < 2 {
            return false;
        }

        let mut hash_map = HashMap::<i32, i32>::new();
        hash_map.insert(0, -1);

        let mut cur_sum = 0;
        for i in 0..n {
            let num = nums[i];
            cur_sum += num;

            if k != 0 {
                cur_sum = cur_sum % k;
            }
            if hash_map.contains_key(&cur_sum) {
                if i as i32 - hash_map.get(&cur_sum).unwrap() > 1 {
                    return true;
                }
            } else {
                hash_map.insert(cur_sum, i as i32);
            }
        }

        false
    }
}
