use std::collections::HashMap;

fn main() {
    println!("{}", Solution::tuple_same_product(vec![2, 3, 4, 6]));
    println!("{}", Solution::tuple_same_product(vec![1, 2, 4, 5, 10]));
    println!(
        "{}",
        Solution::tuple_same_product(vec![1, 2, 3, 4, 6, 8, 12, 24])
    );
}

struct Solution;

impl Solution {
    pub fn tuple_same_product(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut hash_map = HashMap::new();
        let mut res = 0;

        for i in 0..n {
            for j in i + 1..n {
                let product = nums[i] * nums[j];
                res += hash_map.get(&product).unwrap_or(&0) * 8;
                hash_map.entry(product).and_modify(|k| *k += 1).or_insert(1);
            }
        }

        res
    }
}
