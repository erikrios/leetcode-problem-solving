use std::collections::HashMap;

fn main() {
    println!(
        "{:?}",
        Solution::find_anagrams("cbaebabacd".to_string(), "abc".to_string())
    );
    println!(
        "{:?}",
        Solution::find_anagrams("abab".to_string(), "ab".to_string())
    );
}

struct Solution;

impl Solution {
    pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
        let mut mapper = HashMap::new();

        for ch in p.chars() {
            mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1);
        }

        let s: Vec<char> = s.chars().collect();

        let mut window_start = 0;
        let mut ch_mapper = HashMap::new();
        let mut results = Vec::new();

        for window_end in 0..s.len() {
            let ch = s[window_end];
            ch_mapper.entry(ch).and_modify(|e| *e += 1).or_insert(1);

            if window_end >= p.len() - 1 {
                let mut is_anagram = true;
                for (k, &v) in &mapper {
                    if v != *ch_mapper.get(k).unwrap_or(&0) {
                        is_anagram = false;
                        break;
                    }
                }
                if is_anagram {
                    results.push(window_start as i32);
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

        results
    }
}
