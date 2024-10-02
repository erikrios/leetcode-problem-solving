use std::collections::BTreeMap;

fn main() {
    println!("{:?}", Solution::array_rank_transform(vec![40, 10, 20, 30]));
    println!("{:?}", Solution::array_rank_transform(vec![100, 100, 100]));
    println!(
        "{:?}",
        Solution::array_rank_transform(vec![37, 12, 28, 9, 100, 56, 80, 5, 12])
    );
}

struct Solution;

impl Solution {
    pub fn array_rank_transform(arr: Vec<i32>) -> Vec<i32> {
        let mut b_tree_map = BTreeMap::new();

        for &num in &arr {
            b_tree_map.insert(num, 0_usize);
        }

        let keys: Vec<_> = b_tree_map.keys().cloned().collect();

        for (i, k) in keys.into_iter().enumerate() {
            b_tree_map.insert(k, i + 1);
        }

        let mut arr = arr;
        let n = arr.len();

        for num in arr.iter_mut().take(n) {
            *num = *b_tree_map.get(num).unwrap() as i32;
        }

        arr
    }
}
