use std::{
    cmp::Reverse,
    collections::{BinaryHeap, HashSet},
};

fn main() {
    println!("{}", Solution::nth_ugly_number(10));
    println!("{}", Solution::nth_ugly_number(1));
    println!("{}", Solution::nth_ugly_number(1690));
}

struct Solution;

impl Solution {
    pub fn nth_ugly_number(n: i32) -> i32 {
        let mut heap = BinaryHeap::new();
        let mut set = HashSet::new();

        heap.push(Reverse(1));
        set.insert(1);

        let factors = [2, 3, 5];
        let mut current: i32 = 1;

        for _ in 0..n {
            current = heap.pop().unwrap().0;

            for &factor in &factors {
                if let Some(new_ugly) = current.checked_mul(factor) {
                    if set.insert(new_ugly) {
                        heap.push(Reverse(new_ugly));
                    }
                }
            }
        }

        current
    }
}
