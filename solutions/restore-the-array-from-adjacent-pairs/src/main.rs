use std::collections::{HashMap, HashSet};

fn main() {
    println!(
        "{:?}",
        Solution::restore_array(vec![vec![2, 1], vec![3, 4], vec![3, 2]])
    );
    println!(
        "{:?}",
        Solution::restore_array(vec![vec![4, -2], vec![1, 4], vec![-3, 1]])
    );
    println!("{:?}", Solution::restore_array(vec![vec![100000, -100000]]));
}

struct Solution;

impl Solution {
    pub fn restore_array(adjacent_pairs: Vec<Vec<i32>>) -> Vec<i32> {
        let mut freqs = HashMap::new();
        let mut adj_mapper = HashMap::new();

        for adjacent_pair in adjacent_pairs {
            let first_num = adjacent_pair[0];
            let second_num = adjacent_pair[1];

            freqs
                .entry(first_num)
                .and_modify(|freq| *freq = false)
                .or_insert(true);

            freqs
                .entry(second_num)
                .and_modify(|freq| *freq = false)
                .or_insert(true);

            adj_mapper
                .entry(first_num)
                .and_modify(|adjacents: &mut Vec<i32>| adjacents.push(second_num))
                .or_insert(vec![second_num]);

            adj_mapper
                .entry(second_num)
                .and_modify(|adjacents: &mut Vec<i32>| adjacents.push(first_num))
                .or_insert(vec![first_num]);
        }

        let mut element = freqs.into_iter().find(|x| x.1).unwrap().0;
        let mut results = vec![element];
        let mut inserted = HashSet::new();
        inserted.insert(element);

        'outer: loop {
            let adjacents = adj_mapper.get(&element).unwrap();
            for &adj in adjacents {
                if inserted.contains(&adj) {
                    if adjacents.len() == 1 {
                        break 'outer;
                    }
                    continue;
                }

                results.push(adj);
                inserted.insert(adj);
                element = adj;
                continue 'outer;
            }
        }

        results
    }
}
