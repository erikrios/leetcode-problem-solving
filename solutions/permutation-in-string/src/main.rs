use std::collections::HashMap;

fn main() {
    println!(
        "{}",
        Solution::check_inclusion("ab".to_string(), "eidbaooo".to_string())
    );
    println!(
        "{}",
        Solution::check_inclusion("ab".to_string(), "eidboaoo".to_string())
    );
    println!(
        "{}",
        Solution::check_inclusion("ab".to_string(), "ab".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        let mut mapper = HashMap::new();

        for ch in s1.chars() {
            mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1);
        }

        let s: Vec<char> = s2.chars().collect();

        let mut window_start = 0;
        let mut ch_mapper = HashMap::new();

        for window_end in 0..s.len() {
            let ch = s[window_end];
            ch_mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1);

            if window_end >= s1.len() - 1 {
                let mut is_permutation = true;
                for (k, &v) in &mapper {
                    if v != *ch_mapper.get(k).unwrap_or(&0) {
                        is_permutation = false;
                        break;
                    }
                }
                if is_permutation {
                    return true;
                }

                let ch_start = s[window_start];
                if let Some(v) = ch_mapper.get_mut(&ch_start) {
                    if *v == 1 {
                        ch_mapper.remove(&ch_start);
                    } else {
                        *v -= 1;
                    }
                }
                window_start += 1;
            }
        }

        false
    }
}
