use std::collections::BTreeMap;

fn main() {
    println!(
        "{}",
        Solution::max_distance(vec![vec![1, 2, 3], vec![4, 5], vec![1, 2, 3]])
    );
    println!("{}", Solution::max_distance(vec![vec![1], vec![1],]));
    println!(
        "{}",
        Solution::max_distance(vec![vec![1, 2, 3, 4, 5, 6], vec![4, 5], vec![2, 3]])
    );
    println!("{}", Solution::max_distance(vec![vec![1, 5], vec![3, 4],]));
}

struct Solution;

impl Solution {
    pub fn max_distance(arrays: Vec<Vec<i32>>) -> i32 {
        let mut map = BTreeMap::new();

        for (i, array) in arrays.iter().enumerate() {
            let min = array[0];
            let max = array[array.len() - 1];

            map.entry(min)
                .and_modify(|x: &mut Vec<usize>| x.push(i))
                .or_insert(vec![i]);

            map.entry(max)
                .and_modify(|x: &mut Vec<usize>| x.push(i))
                .or_insert(vec![i]);
        }

        struct Str {
            k: i32,
            v: Vec<usize>,
        }

        let mut strs = Vec::with_capacity(map.len());

        for (k, v) in map {
            strs.push(Str { k, v })
        }

        for i in 0..strs.len() / 2 {
            let fwd = &strs[i];
            let bwd = &strs[strs.len() - 1 - i];

            if fwd.v.len() == 1 && bwd.v.len() == 1 && fwd.v[0] == bwd.v[0] {
                let mut max = i32::MIN;

                if i + 1 < strs.len() {
                    let next_fwd = &strs[i + 1];
                    let res = (bwd.k - next_fwd.k).abs();
                    if res > max {
                        max = res;
                    }
                }
                if i as i32 >= 0 {
                    let prev_bwd = &strs[strs.len() - 1 - i - 1];
                    let res = (prev_bwd.k - fwd.k).abs();
                    if res > max {
                        max = res;
                    }
                }

                if max == i32::MIN {
                    continue;
                }

                return max;
            }

            return (bwd.k - fwd.k).abs();
        }

        0
    }
}
