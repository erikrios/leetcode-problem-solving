use std::collections::HashSet;

fn main() {
    println!(
        "{:?}",
        Solution::has_increasing_subarrays(vec![2, 5, 7, 8, 9, 2, 3, 4, 3, 1], 3)
    );
    println!(
        "{:?}",
        Solution::has_increasing_subarrays(vec![1, 2, 3, 4, 4, 4, 4, 5, 6, 7], 5)
    );
    println!("{:?}", Solution::has_increasing_subarrays(vec![-15, 19], 1));
}

struct Solution;

impl Solution {
    pub fn has_increasing_subarrays(nums: Vec<i32>, k: i32) -> bool {
        if k == 1 {
            return true;
        }

        let n = nums.len();
        let mut total_increasing = 0;
        let mut starting_index_increasing_tracker = HashSet::new();

        for i in 0..n {
            if i == 0 {
                continue;
            }

            if i >= k as usize {
                let prev_starter_rem = nums[i - k as usize];
                let next_starter_rem = nums[i - k as usize + 1];
                if next_starter_rem > prev_starter_rem {
                    total_increasing -= 1;
                }
            }

            let cur = nums[i];
            let prev = nums[i - 1];

            if cur > prev {
                total_increasing += 1;
                if total_increasing == k - 1 {
                    let starting_index = (i as i32 - k + 1) as usize;
                    if starting_index >= k as usize
                        && starting_index_increasing_tracker
                            .contains(&(starting_index - k as usize))
                    {
                        return true;
                    }
                    starting_index_increasing_tracker.insert(starting_index);
                }
            }
        }

        false
    }
}
