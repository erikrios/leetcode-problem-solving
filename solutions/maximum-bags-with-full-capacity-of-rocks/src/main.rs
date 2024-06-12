use std::collections::BTreeMap;

fn main() {
    println!(
        "{}",
        Solution::maximum_bags(vec![2, 3, 4, 5], vec![1, 2, 4, 4], 2)
    );
    println!(
        "{}",
        Solution::maximum_bags(vec![10, 2, 2], vec![2, 2, 0], 100)
    );
}

struct Solution;

impl Solution {
    pub fn maximum_bags(capacity: Vec<i32>, rocks: Vec<i32>, additional_rocks: i32) -> i32 {
        let n = capacity.len();
        let mut tree_map = BTreeMap::<i32, Vec<usize>>::new();

        let mut counter = 0;
        for i in 0..n {
            let cap = capacity[i];
            let rock = rocks[i];

            if cap == rock {
                counter += 1;
                continue;
            }

            let key = cap - rock;
            match tree_map.get_mut(&key) {
                Some(val) => {
                    val.push(i);
                }
                None => {
                    tree_map.insert(key, vec![i]);
                }
            }
        }

        let mut additional_rocks = additional_rocks;

        'loops: for (k, vals) in tree_map {
            for _ in vals {
                if additional_rocks - k < 0 {
                    break 'loops;
                }
                additional_rocks -= k;
                counter += 1;
            }
        }

        counter
    }
}
