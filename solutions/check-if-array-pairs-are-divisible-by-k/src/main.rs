use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::can_arrange(vec![1, 2, 3, 4, 5, 10, 6, 7, 8, 9], 5)
    );
    println!("{}", Solution::can_arrange(vec![1, 2, 3, 4, 5, 6], 7));
    println!("{}", Solution::can_arrange(vec![1, 2, 3, 4, 5, 6], 10));
    println!(
        "{}",
        Solution::can_arrange(vec![-1, 1, -2, 2, -3, 3, -4, 4], 3)
    );
    println!(
        "{}",
        Solution::can_arrange(vec![-4, -7, 5, 2, 9, 1, 10, 4, -8, -3], 3)
    );
}

struct Solution;

impl Solution {
    pub fn can_arrange(arr: Vec<i32>, k: i32) -> bool {
        let mut remainder_count = HashMap::new();

        for &num in &arr {
            let mod_res = ((num % k) + k) % k;
            *remainder_count.entry(mod_res).or_insert(0) += 1;
        }

        for (&remainder, &count) in remainder_count.iter() {
            if remainder == 0 {
                if count % 2 != 0 {
                    return false;
                }
            } else {
                let complement = k - remainder;
                if let Some(&complement_count) = remainder_count.get(&complement) {
                    if count != complement_count {
                        return false;
                    }
                } else {
                    return false;
                }
            }
        }

        true
    }
}
