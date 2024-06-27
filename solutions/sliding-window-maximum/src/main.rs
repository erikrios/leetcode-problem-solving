use std::collections::BTreeMap;

fn main() {
    println!(
        "{:?}",
        Solution::max_sliding_window(vec![1, 3, -1, -3, 5, 3, 6, 7], 3)
    );
    println!("{:?}", Solution::max_sliding_window(vec![1], 1));
}

struct Solution;

impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut window_start = 0;
        let mut tree_map = BTreeMap::new();
        let mut results = Vec::with_capacity(nums.len() - k as usize + 1);

        for (window_end, num) in nums.iter().enumerate() {
            tree_map.entry(num).and_modify(|e| *e += 1).or_insert(1);
            if window_end >= k as usize - 1 {
                let &max = tree_map.keys().next_back().unwrap();
                results.push(*max);
                let num = nums[window_start];
                if let Some(count) = tree_map.get_mut(&num) {
                    if *count == 1 {
                        tree_map.remove(&num);
                    } else {
                        *count -= 1;
                    }
                }
                window_start += 1;
            }
        }

        results
    }
}
