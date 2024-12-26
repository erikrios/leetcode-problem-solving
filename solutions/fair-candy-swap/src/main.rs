use std::collections::HashSet;

fn main() {
    println!("{:?}", Solution::fair_candy_swap(vec![1, 1], vec![2, 2]));
    println!("{:?}", Solution::fair_candy_swap(vec![1, 2], vec![2, 3]));
    println!("{:?}", Solution::fair_candy_swap(vec![2], vec![1, 3]));
}

struct Solution;

impl Solution {
    pub fn fair_candy_swap(alice_sizes: Vec<i32>, bob_sizes: Vec<i32>) -> Vec<i32> {
        let alice_total: i32 = alice_sizes.iter().sum();
        let bob_total: i32 = bob_sizes.iter().sum();
        let total = alice_total + bob_total;
        let each = total / 2;

        let bob_hash_set: HashSet<i32> = HashSet::from_iter(bob_sizes);

        for num in alice_sizes {
            let rem = alice_total - num;
            let target = each - rem;
            if bob_hash_set.contains(&target) {
                return vec![num, target];
            }
        }

        vec![]
    }
}
