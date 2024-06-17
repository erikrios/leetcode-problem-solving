use std::collections::HashMap;

fn main() {
    println!(
        "{:?}",
        Solution::partition_labels("ababcbacadefegdehijhklij".to_string())
    );
    println!("{:?}", Solution::partition_labels("eccbbbbdec".to_string()));
}

struct Solution;

impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let mut mapper = HashMap::<char, usize>::new();
        let mut results = Vec::new();

        let s = s.chars();

        for character in s {
            match mapper.get(&character) {
                Some(&v) => {
                    if v == results.len() {
                        let last_index = results.len() - 1;
                        results[last_index] += 1;
                    } else {
                        let need_updates: Vec<(char, usize)> = mapper
                            .iter()
                            .filter(|(_, &val)| val > v)
                            .map(|(k, _)| (k.clone(), v))
                            .collect();
                        for (k, _) in need_updates {
                            mapper.insert(k, v);
                        }

                        let sum: i32 = results.iter().skip(v).copied().sum();

                        results.truncate(v);
                        results[v - 1] += sum + 1;
                    }
                }
                None => {
                    mapper.insert(character, results.len() + 1);
                    results.push(1);
                }
            }
        }

        results
    }
}
