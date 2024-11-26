use std::collections::HashSet;

fn main() {
    println!(
        "{}",
        Solution::find_champion(3, vec![vec![0, 1], vec![1, 2]])
    );
    println!(
        "{}",
        Solution::find_champion(4, vec![vec![0, 2], vec![1, 3], vec![1, 2]])
    );
}

struct Solution;

impl Solution {
    pub fn find_champion(n: i32, edges: Vec<Vec<i32>>) -> i32 {
        let mut hash_set = HashSet::new();

        for i in 0..n {
            hash_set.insert(i);
        }

        for edge in edges {
            hash_set.remove(&edge[1]);
        }

        if hash_set.len() > 1 {
            return -1;
        }

        hash_set.into_iter().collect::<Vec<i32>>()[0]
    }
}
